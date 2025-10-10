package service

import (
	"context"

	"github.com/google/uuid"
	"task-management/model/web"
)

// ProjectService mendefinisikan kontrak bisnis logic
// untuk operasi CRUD pada entitas Project.
type ProjectService interface {
	// Create membuat project baru berdasarkan data yang diberikan.
	Create(ctx context.Context, request web.ProjectCreateRequest) web.ProjectResponse

	// Update memperbarui project berdasarkan ID yang diberikan.
	Update(ctx context.Context, request web.ProjectUpdateRequest) web.ProjectResponse

	// Delete menghapus project berdasarkan ID yang diberikan.
	Delete(ctx context.Context, projectId uuid.UUID)

	// FindById mengambil data project berdasarkan ID-nya.
	FindById(ctx context.Context, projectId uuid.UUID) web.ProjectResponse

	// FindByUserId mengambil semua project yang dimiliki oleh user tertentu.
	FindByUserId(ctx context.Context, userId uuid.UUID) []web.ProjectResponse

	// FindAll mengambil semua project yang ada di sistem.
	FindAll(ctx context.Context) []web.ProjectResponse
}
