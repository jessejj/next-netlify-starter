package handlers_test

import (
	"context"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/getfider/fider/app"
	"github.com/getfider/fider/app/models/entity"
	"github.com/getfider/fider/app/models/query"
	"github.com/getfider/fider/app/pkg/bus"
	"github.com/getfider/fider/app/services/httpclient"

	"github.com/getfider/fider/app/pkg/mock"

	"github.com/getfider/fider/app/handlers"
	. "github.com/getfider/fider/app/pkg/assert"
)

func TestGravatarHandler(t *testing.T) {
	RegisterT(t)
	bus.Init(httpclient.Service{})

	server := mock.NewServer()

	user := &entity.User{
		ID:     3,
		Name:   "Darth Vader",
		Email:  "darthvader.fider@gmail.com",
		Tenant: mock.DemoTenant,
	}

	bus.AddHandler(func(ctx context.Context, q *query.GetUserByID) error {
		if q.UserID == user.ID {
			q.Result = user
			return nil
		}
		return app.ErrNotFound
	})

	code, response := server.
		WithURL("https://demo.test.fider.io/?size=50").
		OnTenant(mock.DemoTenant).
		AddParam("id", user.ID).
		AddParam("name", user.Name).
		Execute(handlers.Gravatar())

	expectedAvatar := []byte{255, 216, 255, 224, 0, 16, 74, 70, 73, 70, 0, 1, 1, 1, 0, 96, 0, 96, 0, 0, 255, 254, 0, 59, 67, 82, 69, 65, 84, 79, 82, 58, 32, 103, 100, 45, 106, 112, 101, 103, 32, 118, 49, 46, 48, 32, 40, 117, 115, 105, 110, 103, 32, 73, 74, 71, 32, 74, 80, 69, 71, 32, 118, 54, 50, 41, 44, 32, 113, 117, 97, 108, 105, 116, 121, 32, 61, 32, 57, 48, 10, 255, 219, 0, 67, 0, 3, 2, 2, 3, 2, 2, 3, 3, 3, 3, 4, 3, 3, 4, 5, 8, 5, 5, 4, 4, 5, 10, 7, 7, 6, 8, 12, 10, 12, 12, 11, 10, 11, 11, 13, 14, 18, 16, 13, 14, 17, 14, 11, 11, 16, 22, 16, 17, 19, 20, 21, 21, 21, 12, 15, 23, 24, 22, 20, 24, 18, 20, 21, 20, 255, 219, 0, 67, 1, 3, 4, 4, 5, 4, 5, 9, 5, 5, 9, 20, 13, 11, 13, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 255, 192, 0, 17, 8, 0, 50, 0, 50, 3, 1, 34, 0, 2, 17, 1, 3, 17, 1, 255, 196, 0, 31, 0, 0, 1, 5, 1, 1, 1, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 255, 196, 0, 181, 16, 0, 2, 1, 3, 3, 2, 4, 3, 5, 5, 4, 4, 0, 0, 1, 125, 1, 2, 3, 0, 4, 17, 5, 18, 33, 49, 65, 6, 19, 81, 97, 7, 34, 113, 20, 50, 129, 145, 161, 8, 35, 66, 177, 193, 21, 82, 209, 240, 36, 51, 98, 114, 130, 9, 10, 22, 23, 24, 25, 26, 37, 38, 39, 40, 41, 42, 52, 53, 54, 55, 56, 57, 58, 67, 68, 69, 70, 71, 72, 73, 74, 83, 84, 85, 86, 87, 88, 89, 90, 99, 100, 101, 102, 103, 104, 105, 106, 115, 116, 117, 118, 119, 120, 121, 122, 131, 132, 133, 134, 135, 136, 137, 138, 146, 147, 148, 149, 150, 151, 152, 153, 154, 162, 163, 164, 165, 166, 167, 168, 169, 170, 178, 179, 180, 181, 182, 183, 184, 185, 186, 194, 195, 196, 197, 198, 199, 200, 201, 202, 210, 211, 212, 213, 214, 215, 216, 217, 218, 225, 226, 227, 228, 229, 230, 231, 232, 233, 234, 241, 242, 243, 244, 245, 246, 247, 248, 249, 250, 255, 196, 0, 31, 1, 0, 3, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 0, 0, 0, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 255, 196, 0, 181, 17, 0, 2, 1, 2, 4, 4, 3, 4, 7, 5, 4, 4, 0, 1, 2, 119, 0, 1, 2, 3, 17, 4, 5, 33, 49, 6, 18, 65, 81, 7, 97, 113, 19, 34, 50, 129, 8, 20, 66, 145, 161, 177, 193, 9, 35, 51, 82, 240, 21, 98, 114, 209, 10, 22, 36, 52, 225, 37, 241, 23, 24, 25, 26, 38, 39, 40, 41, 42, 53, 54, 55, 56, 57, 58, 67, 68, 69, 70, 71, 72, 73, 74, 83, 84, 85, 86, 87, 88, 89, 90, 99, 100, 101, 102, 103, 104, 105, 106, 115, 116, 117, 118, 119, 120, 121, 122, 130, 131, 132, 133, 134, 135, 136, 137, 138, 146, 147, 148, 149, 150, 151, 152, 153, 154, 162, 163, 164, 165, 166, 167, 168, 169, 170, 178, 179, 180, 181, 182, 183, 184, 185, 186, 194, 195, 196, 197, 198, 199, 200, 201, 202, 210, 211, 212, 213, 214, 215, 216, 217, 218, 226, 227, 228, 229, 230, 231, 232, 233, 234, 242, 243, 244, 245, 246, 247, 248, 249, 250, 255, 218, 0, 12, 3, 1, 0, 2, 17, 3, 17, 0, 63, 0, 252, 207, 109, 22, 21, 199, 12, 61, 112, 213, 57, 209, 237, 33, 182, 105, 165, 46, 168, 163, 183, 57, 62, 149, 245, 7, 236, 5, 225, 255, 0, 3, 235, 159, 181, 46, 155, 99, 227, 221, 31, 79, 214, 244, 137, 244, 107, 150, 182, 179, 213, 33, 89, 160, 123, 133, 101, 32, 148, 111, 148, 144, 130, 67, 200, 237, 94, 239, 255, 0, 5, 55, 248, 123, 240, 166, 203, 194, 254, 23, 210, 126, 29, 120, 111, 194, 158, 24, 214, 165, 214, 86, 43, 171, 171, 8, 161, 179, 45, 7, 145, 35, 54, 66, 224, 20, 12, 83, 45, 142, 14, 209, 223, 7, 179, 151, 70, 212, 79, 38, 51, 189, 175, 61, 207, 139, 79, 193, 22, 210, 244, 49, 168, 222, 70, 100, 114, 155, 150, 44, 0, 188, 142, 6, 113, 147, 214, 188, 207, 84, 125, 62, 210, 249, 173, 96, 178, 44, 85, 182, 231, 130, 73, 252, 171, 236, 77, 122, 230, 211, 83, 248, 123, 13, 181, 133, 237, 189, 237, 205, 148, 40, 151, 83, 7, 85, 141, 36, 85, 220, 9, 98, 113, 180, 237, 206, 115, 192, 53, 243, 111, 135, 188, 51, 164, 71, 227, 107, 121, 117, 45, 95, 79, 186, 182, 50, 3, 32, 176, 147, 206, 216, 51, 215, 177, 35, 220, 10, 43, 69, 69, 46, 84, 42, 18, 114, 187, 155, 216, 169, 160, 248, 58, 195, 92, 133, 198, 207, 42, 120, 240, 94, 41, 19, 4, 3, 208, 253, 56, 171, 147, 124, 54, 180, 83, 247, 98, 3, 25, 232, 223, 227, 95, 74, 254, 206, 158, 16, 240, 47, 141, 255, 0, 110, 47, 15, 120, 71, 80, 180, 77, 127, 194, 122, 134, 149, 112, 37, 242, 103, 150, 5, 44, 45, 222, 85, 108, 198, 202, 195, 13, 26, 140, 103, 191, 61, 13, 125, 161, 241, 7, 246, 40, 248, 19, 166, 207, 226, 11, 75, 123, 61, 86, 11, 230, 137, 26, 22, 143, 85, 148, 165, 129, 59, 137, 3, 113, 57, 227, 105, 195, 238, 60, 142, 153, 171, 167, 21, 39, 203, 107, 179, 58, 142, 80, 138, 159, 53, 147, 63, 36, 207, 195, 155, 44, 253, 216, 255, 0, 241, 239, 254, 42, 138, 250, 238, 251, 224, 231, 192, 219, 11, 219, 139, 89, 126, 37, 220, 44, 144, 200, 209, 178, 181, 237, 190, 65, 7, 4, 31, 147, 218, 138, 233, 246, 17, 236, 190, 244, 115, 123, 106, 157, 223, 226, 121, 191, 236, 225, 225, 237, 38, 235, 196, 90, 30, 175, 127, 4, 114, 125, 140, 93, 100, 191, 241, 6, 27, 54, 253, 48, 77, 124, 243, 241, 167, 198, 35, 226, 39, 196, 237, 94, 235, 78, 139, 26, 104, 185, 107, 125, 58, 218, 17, 144, 33, 86, 33, 118, 129, 221, 190, 241, 247, 106, 244, 107, 47, 23, 73, 225, 15, 133, 90, 133, 196, 13, 229, 220, 75, 20, 144, 68, 71, 4, 51, 182, 220, 143, 112, 9, 63, 133, 120, 5, 178, 204, 242, 55, 144, 28, 176, 82, 78, 204, 228, 14, 245, 203, 94, 126, 226, 130, 245, 59, 240, 112, 213, 212, 126, 136, 244, 13, 55, 74, 178, 240, 239, 135, 124, 77, 160, 235, 186, 69, 235, 248, 154, 254, 218, 6, 210, 86, 221, 247, 164, 18, 137, 163, 102, 50, 42, 57, 249, 140, 94, 114, 224, 130, 65, 113, 144, 59, 114, 87, 90, 123, 216, 216, 219, 93, 164, 166, 11, 168, 159, 99, 194, 210, 1, 42, 54, 73, 4, 40, 59, 128, 192, 28, 144, 57, 53, 235, 31, 7, 143, 143, 126, 30, 106, 54, 94, 32, 240, 236, 210, 232, 154, 157, 131, 75, 50, 203, 106, 26, 43, 150, 67, 20, 76, 81, 164, 94, 118, 178, 49, 194, 156, 242, 132, 17, 212, 87, 97, 241, 150, 255, 0, 196, 30, 54, 240, 54, 163, 172, 248, 150, 120, 237, 158, 105, 197, 202, 202, 2, 192, 215, 239, 130, 0, 242, 219, 27, 74, 156, 238, 0, 99, 43, 129, 201, 192, 227, 183, 83, 177, 212, 229, 146, 91, 220, 227, 127, 103, 175, 137, 175, 225, 175, 142, 190, 1, 241, 13, 228, 134, 40, 99, 186, 58, 117, 212, 209, 49, 70, 141, 103, 86, 133, 159, 35, 160, 11, 38, 239, 248, 11, 12, 98, 189, 215, 246, 171, 241, 135, 142, 45, 245, 173, 91, 77, 209, 124, 79, 126, 44, 245, 139, 159, 54, 243, 124, 196, 201, 147, 144, 64, 110, 170, 164, 99, 229, 30, 149, 241, 45, 148, 139, 230, 52, 82, 62, 200, 228, 227, 118, 126, 235, 127, 11, 126, 127, 161, 53, 245, 126, 187, 227, 120, 254, 34, 120, 99, 70, 241, 35, 184, 150, 226, 242, 216, 45, 216, 207, 41, 118, 152, 89, 129, 244, 203, 124, 224, 127, 118, 68, 174, 202, 18, 188, 92, 14, 76, 92, 57, 92, 106, 37, 177, 226, 171, 240, 158, 2, 163, 204, 190, 99, 39, 241, 125, 123, 209, 91, 242, 106, 110, 36, 108, 103, 25, 52, 86, 156, 144, 254, 83, 139, 219, 213, 254, 98, 75, 109, 16, 120, 175, 192, 113, 105, 198, 65, 19, 191, 207, 28, 135, 160, 96, 73, 25, 246, 237, 248, 215, 157, 234, 191, 12, 124, 83, 160, 193, 113, 52, 154, 116, 173, 104, 138, 124, 201, 237, 220, 58, 21, 235, 147, 131, 211, 128, 121, 244, 174, 183, 71, 241, 141, 151, 135, 60, 63, 100, 247, 83, 22, 148, 198, 25, 97, 143, 150, 111, 240, 252, 107, 152, 241, 151, 197, 61, 91, 197, 182, 198, 199, 119, 216, 244, 204, 130, 109, 227, 57, 47, 142, 155, 219, 191, 211, 165, 101, 87, 217, 181, 119, 185, 215, 134, 85, 163, 39, 20, 189, 219, 255, 0, 86, 56, 239, 51, 228, 219, 142, 125, 107, 214, 62, 16, 252, 25, 241, 47, 199, 75, 235, 75, 107, 101, 187, 155, 78, 179, 30, 76, 183, 140, 75, 45, 180, 99, 36, 12, 177, 194, 168, 207, 78, 131, 53, 15, 134, 252, 1, 168, 63, 129, 228, 241, 11, 91, 88, 106, 0, 198, 124, 184, 46, 98, 102, 109, 163, 142, 10, 144, 115, 199, 124, 213, 95, 17, 106, 30, 37, 248, 117, 161, 157, 30, 211, 84, 159, 79, 178, 212, 80, 61, 205, 189, 155, 180, 105, 38, 224, 172, 59, 231, 161, 0, 250, 226, 185, 121, 26, 179, 107, 67, 181, 212, 231, 188, 96, 245, 48, 190, 38, 248, 13, 62, 29, 120, 170, 243, 71, 77, 98, 203, 89, 16, 57, 95, 58, 206, 64, 227, 232, 112, 72, 4, 122, 100, 214, 175, 195, 127, 19, 182, 157, 165, 221, 233, 247, 26, 149, 188, 86, 79, 47, 157, 29, 172, 163, 14, 36, 192, 5, 149, 136, 192, 4, 0, 8, 207, 59, 87, 142, 5, 112, 166, 202, 115, 100, 110, 202, 17, 0, 144, 71, 188, 247, 98, 9, 199, 228, 42, 189, 17, 151, 44, 174, 141, 37, 14, 120, 114, 201, 158, 181, 37, 234, 51, 177, 15, 193, 39, 24, 52, 87, 147, 110, 62, 166, 138, 232, 250, 199, 145, 195, 245, 37, 252, 223, 128, 249, 216, 179, 140, 146, 112, 160, 115, 244, 162, 220, 3, 113, 16, 35, 32, 176, 227, 241, 162, 138, 228, 61, 37, 178, 62, 222, 240, 253, 180, 48, 120, 14, 241, 34, 137, 35, 65, 230, 0, 168, 160, 1, 197, 121, 119, 199, 187, 120, 158, 227, 74, 221, 26, 55, 250, 29, 185, 229, 71, 247, 18, 138, 43, 210, 173, 252, 63, 184, 249, 236, 63, 241, 126, 103, 156, 120, 158, 24, 227, 248, 96, 74, 162, 169, 254, 215, 65, 144, 49, 199, 144, 213, 230, 180, 81, 94, 124, 186, 30, 229, 47, 132, 40, 162, 138, 147, 99, 255, 217}
	Expect(code).Equals(http.StatusOK)
	Expect(response.Body.Bytes()).Equals(expectedAvatar)
}

func TestGravatarNotFound_LetterAvatarHandler(t *testing.T) {
	RegisterT(t)
	bus.Init(httpclient.Service{})

	server := mock.NewServer()

	user := &entity.User{
		ID:     3,
		Name:   "Darth Vader",
		Email:  "darthvader.fider1234567890@gmail.com",
		Tenant: mock.DemoTenant,
	}

	bus.AddHandler(func(ctx context.Context, q *query.GetUserByID) error {
		if q.UserID == user.ID {
			q.Result = user
			return nil
		}
		return app.ErrNotFound
	})

	code, response := server.
		OnTenant(mock.DemoTenant).
		WithURL("https://demo.test.fider.io/?size=50").
		AddParam("id", user.ID).
		AddParam("name", user.Name).
		Execute(handlers.Gravatar())

	expectedAvatar := []byte{137, 80, 78, 71, 13, 10, 26, 10, 0, 0, 0, 13, 73, 72, 68, 82, 0, 0, 0, 50, 0, 0, 0, 50, 8, 2, 0, 0, 0, 145, 93, 31, 230, 0, 0, 3, 6, 73, 68, 65, 84, 120, 156, 236, 152, 75, 79, 19, 81, 20, 128, 207, 204, 212, 210, 18, 90, 232, 211, 34, 196, 40, 6, 65, 129, 0, 209, 2, 11, 100, 97, 140, 113, 225, 2, 127, 3, 198, 196, 184, 55, 38, 38, 36, 198, 132, 168, 11, 149, 152, 24, 35, 97, 99, 220, 24, 23, 46, 140, 186, 49, 38, 242, 42, 20, 10, 72, 241, 149, 64, 75, 12, 125, 208, 86, 104, 7, 90, 103, 134, 206, 29, 211, 14, 25, 154, 66, 116, 46, 96, 101, 113, 191, 213, 225, 204, 153, 123, 191, 204, 185, 15, 82, 205, 197, 87, 157, 176, 255, 160, 255, 183, 192, 246, 16, 45, 28, 136, 22, 14, 68, 11, 7, 162, 133, 3, 209, 194, 129, 104, 225, 64, 180, 112, 32, 90, 56, 104, 212, 20, 117, 84, 158, 57, 119, 248, 172, 242, 103, 50, 157, 138, 164, 34, 19, 97, 143, 55, 54, 171, 36, 105, 138, 190, 209, 114, 93, 199, 20, 45, 115, 203, 15, 39, 31, 109, 29, 228, 90, 211, 213, 131, 197, 118, 1, 173, 223, 113, 223, 75, 163, 244, 30, 104, 89, 245, 150, 26, 243, 241, 188, 228, 133, 35, 231, 125, 9, 255, 125, 79, 111, 52, 21, 5, 0, 36, 161, 53, 97, 181, 177, 162, 1, 0, 94, 124, 127, 25, 74, 134, 243, 70, 232, 168, 108, 7, 0, 119, 120, 226, 175, 78, 216, 77, 124, 58, 211, 255, 192, 211, 251, 120, 250, 201, 251, 31, 31, 214, 81, 186, 170, 244, 104, 79, 251, 109, 147, 206, 36, 63, 29, 12, 12, 203, 65, 75, 121, 75, 222, 139, 173, 229, 173, 27, 53, 139, 67, 106, 38, 194, 211, 154, 138, 76, 143, 134, 220, 31, 23, 7, 251, 102, 250, 187, 71, 110, 241, 34, 111, 212, 26, 174, 52, 116, 201, 79, 63, 69, 103, 18, 60, 155, 145, 112, 56, 243, 94, 108, 203, 138, 38, 215, 147, 147, 145, 169, 189, 215, 202, 197, 23, 247, 189, 158, 127, 3, 0, 205, 246, 102, 171, 222, 42, 247, 209, 21, 116, 1, 192, 177, 178, 42, 139, 222, 162, 84, 154, 117, 166, 234, 178, 106, 0, 24, 13, 185, 213, 116, 112, 183, 59, 209, 21, 26, 5, 0, 138, 2, 167, 227, 148, 156, 217, 236, 163, 227, 180, 82, 230, 116, 56, 41, 74, 126, 170, 170, 131, 187, 213, 10, 174, 133, 228, 192, 166, 183, 201, 193, 92, 124, 62, 156, 92, 202, 106, 109, 246, 177, 181, 60, 19, 71, 127, 197, 190, 254, 252, 86, 8, 45, 36, 33, 94, 20, 0, 160, 180, 168, 84, 73, 14, 7, 71, 0, 160, 214, 92, 99, 212, 26, 0, 192, 168, 53, 156, 48, 215, 102, 242, 129, 17, 245, 35, 239, 246, 56, 213, 208, 153, 35, 70, 16, 121, 37, 35, 239, 53, 154, 162, 157, 217, 15, 230, 116, 56, 105, 42, 51, 203, 192, 226, 96, 129, 180, 76, 58, 19, 147, 157, 50, 198, 45, 43, 201, 80, 50, 60, 31, 247, 41, 189, 147, 247, 160, 63, 177, 16, 88, 11, 22, 72, 171, 222, 114, 82, 14, 230, 86, 230, 114, 243, 242, 210, 174, 183, 212, 217, 139, 109, 117, 214, 186, 220, 173, 240, 207, 181, 180, 140, 246, 82, 117, 39, 0, 172, 112, 241, 220, 91, 40, 187, 188, 92, 162, 132, 24, 154, 185, 220, 208, 197, 80, 52, 146, 208, 112, 1, 180, 40, 138, 170, 53, 215, 116, 183, 221, 172, 40, 57, 4, 0, 125, 222, 126, 36, 161, 220, 2, 150, 103, 189, 81, 47, 0, 52, 218, 50, 119, 209, 108, 236, 115, 156, 79, 96, 77, 161, 234, 78, 84, 184, 219, 209, 35, 73, 168, 228, 64, 9, 67, 51, 0, 32, 34, 241, 217, 151, 231, 158, 165, 201, 173, 149, 3, 129, 161, 38, 123, 227, 70, 140, 179, 216, 119, 162, 37, 239, 121, 0, 224, 69, 97, 44, 228, 126, 235, 127, 231, 79, 44, 108, 91, 57, 30, 158, 224, 210, 156, 78, 163, 227, 210, 220, 88, 120, 28, 87, 139, 82, 243, 179, 155, 81, 107, 80, 78, 38, 9, 128, 21, 88, 54, 123, 247, 253, 25, 123, 177, 189, 136, 209, 242, 162, 16, 73, 69, 112, 181, 84, 125, 45, 86, 88, 101, 133, 85, 220, 161, 119, 96, 163, 176, 79, 255, 59, 37, 90, 56, 16, 45, 28, 136, 22, 14, 68, 11, 7, 162, 133, 3, 209, 194, 129, 104, 225, 240, 59, 0, 0, 255, 255, 193, 108, 28, 164, 14, 189, 96, 131, 0, 0, 0, 0, 73, 69, 78, 68, 174, 66, 96, 130}
	Expect(code).Equals(http.StatusOK)
	bytes, _ := ioutil.ReadAll(response.Body)
	Expect(bytes).Equals(expectedAvatar)
}

func TestUnknownUser_LetterAvatarHandler(t *testing.T) {
	RegisterT(t)
	bus.Init(httpclient.Service{})

	server := mock.NewServer()
	code, response := server.
		OnTenant(mock.DemoTenant).
		WithURL("https://demo.test.fider.io/?size=50").
		AddParam("id", 0).
		AddParam("name", "Jon Snow").
		Execute(handlers.Gravatar())

	expectedAvatar := []byte{137, 80, 78, 71, 13, 10, 26, 10, 0, 0, 0, 13, 73, 72, 68, 82, 0, 0, 0, 50, 0, 0, 0, 50, 8, 2, 0, 0, 0, 145, 93, 31, 230, 0, 0, 2, 137, 73, 68, 65, 84, 120, 156, 236, 152, 75, 79, 19, 81, 20, 199, 255, 3, 51, 244, 65, 95, 80, 74, 43, 160, 149, 135, 36, 5, 177, 169, 137, 40, 42, 209, 152, 200, 74, 19, 87, 110, 248, 84, 126, 7, 227, 70, 23, 110, 140, 137, 154, 248, 32, 33, 152, 72, 20, 13, 6, 8, 62, 210, 74, 209, 218, 161, 239, 55, 237, 152, 67, 167, 37, 141, 36, 222, 139, 19, 202, 226, 254, 114, 51, 28, 78, 239, 220, 249, 245, 204, 205, 153, 73, 101, 237, 225, 61, 28, 63, 58, 218, 45, 112, 48, 66, 139, 7, 161, 197, 131, 208, 226, 65, 104, 241, 32, 180, 120, 16, 90, 60, 8, 45, 30, 132, 22, 15, 242, 127, 157, 237, 61, 133, 241, 243, 20, 124, 120, 141, 76, 162, 177, 164, 130, 145, 41, 156, 24, 134, 205, 65, 255, 22, 114, 248, 21, 198, 151, 143, 40, 230, 143, 74, 203, 100, 129, 219, 71, 129, 210, 165, 103, 44, 54, 92, 190, 5, 187, 107, 127, 142, 98, 130, 163, 23, 167, 39, 176, 244, 20, 234, 246, 145, 104, 253, 77, 232, 26, 57, 105, 26, 214, 222, 97, 251, 59, 180, 42, 122, 188, 152, 184, 8, 179, 21, 211, 115, 120, 254, 0, 187, 21, 150, 101, 12, 221, 91, 138, 9, 253, 39, 41, 136, 108, 96, 125, 25, 105, 21, 153, 36, 194, 235, 120, 255, 10, 245, 210, 14, 140, 48, 174, 100, 104, 181, 58, 26, 95, 178, 82, 106, 201, 199, 163, 248, 246, 153, 130, 114, 169, 29, 90, 165, 2, 149, 199, 238, 162, 157, 164, 254, 68, 244, 171, 158, 175, 238, 98, 229, 13, 215, 74, 70, 55, 136, 149, 5, 212, 106, 232, 148, 105, 39, 221, 184, 139, 241, 16, 172, 246, 67, 44, 99, 180, 86, 124, 11, 11, 143, 145, 138, 83, 236, 232, 165, 205, 62, 55, 143, 43, 183, 209, 235, 109, 171, 22, 128, 68, 12, 47, 31, 97, 241, 9, 194, 27, 168, 148, 41, 227, 25, 196, 236, 29, 186, 179, 204, 24, 221, 32, 154, 196, 34, 52, 58, 58, 49, 56, 138, 201, 25, 152, 45, 8, 206, 66, 141, 210, 230, 99, 128, 185, 90, 157, 50, 186, 204, 52, 90, 144, 244, 191, 154, 70, 71, 103, 31, 134, 198, 104, 72, 141, 124, 173, 74, 205, 98, 249, 197, 222, 92, 137, 90, 63, 27, 204, 213, 26, 11, 34, 112, 129, 130, 103, 247, 145, 207, 234, 73, 171, 77, 15, 202, 69, 58, 246, 15, 97, 242, 18, 5, 233, 29, 26, 77, 82, 170, 30, 40, 102, 198, 171, 49, 87, 43, 249, 91, 15, 134, 167, 244, 192, 210, 13, 127, 128, 130, 66, 22, 185, 52, 5, 212, 214, 247, 202, 118, 238, 42, 53, 207, 58, 146, 132, 209, 198, 41, 153, 29, 198, 171, 49, 87, 43, 22, 33, 51, 151, 7, 103, 130, 240, 249, 233, 185, 235, 242, 64, 81, 232, 163, 213, 183, 250, 156, 108, 18, 171, 75, 56, 59, 131, 190, 1, 220, 156, 167, 130, 85, 138, 176, 247, 232, 61, 34, 165, 226, 199, 38, 227, 213, 36, 142, 223, 183, 76, 22, 132, 174, 147, 83, 147, 82, 129, 156, 194, 107, 45, 211, 124, 126, 4, 166, 225, 116, 239, 103, 106, 53, 108, 109, 226, 211, 162, 126, 175, 13, 214, 170, 99, 181, 195, 225, 134, 44, 211, 27, 75, 34, 70, 155, 250, 64, 186, 157, 176, 57, 233, 205, 162, 84, 160, 54, 198, 252, 216, 169, 195, 223, 32, 242, 25, 26, 255, 36, 151, 162, 113, 88, 142, 233, 219, 169, 208, 226, 65, 104, 241, 32, 180, 120, 16, 90, 60, 8, 45, 30, 132, 22, 15, 66, 139, 135, 63, 1, 0, 0, 255, 255, 117, 91, 183, 108, 67, 255, 85, 68, 0, 0, 0, 0, 73, 69, 78, 68, 174, 66, 96, 130}
	Expect(code).Equals(http.StatusOK)
	bytes, _ := ioutil.ReadAll(response.Body)
	Expect(bytes).Equals(expectedAvatar)
}
