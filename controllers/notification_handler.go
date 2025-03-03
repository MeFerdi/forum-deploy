package controllers

import (
	"html/template"
	"log"
	"net/http"

	"forum/utils"
)

type NotificationHandler struct{}

func NewNotificationHandler() *NotificationHandler {
	return &NotificationHandler{}
}

func (nh *NotificationHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		utils.RenderErrorPage(w, http.StatusMethodNotAllowed, utils.ErrMethodNotAllowed)
		return
	}
	nh.handleGetNotifications(w, r)
}

func (nh *NotificationHandler) handleGetNotifications(w http.ResponseWriter, r *http.Request) {
	// Get current user ID from session
	cookie, err := r.Cookie("session_token")
	if err != nil {
		http.Redirect(w, r, "/signin", http.StatusSeeOther)
		return
	}

	userID, err := utils.ValidateSession(utils.GlobalDB, cookie.Value)
	if err != nil {
		http.Redirect(w, r, "/signin", http.StatusSeeOther)
		return
	}

	// Fetch notifications for the user
	notifications, err := nh.getUserNotifications(userID)
	if err != nil {
		log.Printf("Error fetching notifications: %v", err)
		utils.RenderErrorPage(w, http.StatusInternalServerError, utils.ErrInternalServer)
		return
	}

	data := struct {
		Notifications []utils.Notification
		IsLoggedIn    bool
		CurrentUserID string
	}{
		Notifications: notifications,
		IsLoggedIn:    true,
		CurrentUserID: userID,
	}

	tmpl, err := template.ParseFiles("templates/notifications.html")
	if err != nil {
		log.Printf("Error parsing template: %v", err)
		utils.RenderErrorPage(w, http.StatusInternalServerError, utils.ErrTemplateLoad)
		return
	}

	if err := tmpl.Execute(w, data); err != nil {
		log.Printf("Error executing template: %v", err)
		utils.RenderErrorPage(w, http.StatusInternalServerError, utils.ErrTemplateExec)
	}
}

func (nh *NotificationHandler) getUserNotifications(userID string) ([]utils.Notification, error) {
	rows, err := utils.GlobalDB.Query(`
        SELECT n.id, n.type, n.created_at, n.post_id, u.username, u.profile_pic
        FROM notifications n
        JOIN users u ON n.actor_id = u.id
        WHERE n.user_id = ?
        ORDER BY n.created_at DESC
    `, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var notifications []utils.Notification
	for rows.Next() {
		var n utils.Notification
		err := rows.Scan(&n.ID, &n.Type, &n.CreatedAt, &n.PostID, &n.ActorName, &n.ActorProfilePic)
		if err != nil {
			log.Printf("Error scanning notification: %v", err)
			continue
		}
		n.CreatedAtFormatted = FormatTimeAgo(n.CreatedAt)
		notifications = append(notifications, n)
	}
	return notifications, nil
}
