package dashboard

import (
	"context"
	"net/http"

	"github.com/google/uuid"

	"github.com/jqwez/caselore/view/components"
)

func RegisterToast(mux *http.ServeMux) {
	toastMux := http.NewServeMux()
	mux.Handle("/htmx/dashboard/toast/", http.StripPrefix("/htmx/dashboard/toast", toastMux))
	toastMux.HandleFunc("/404", HandleToast404)
}

func HandleToast404(w http.ResponseWriter, r *http.Request) {
	props := components.ToastProps{ID: uuid.New().String(), Text: "Request Invalid: 404", Visible: true}
	components.ToastContainerError(props).Render(context.Background(), w)
}

func ToastSendSuccess(w http.ResponseWriter, msg string) {
	props := components.ToastProps{ID: uuid.New().String(), Text: msg, Visible: true}
	components.ToastContainer(props).Render(context.Background(), w)
}

func HandleToastRemove(w http.ResponseWriter, r *http.Request) {
	props := components.ToastProps{Text: "", Visible: false}
	components.ToastContainer(props).Render(context.Background(), w)
}
