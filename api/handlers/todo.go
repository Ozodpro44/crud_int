package handlers

import "app/models"

func (h *Handlers) CreateTodo() {
	h.storage.GetTodoRepo().CreateTodo(ctx, models.Todo{})
}
