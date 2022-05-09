package handler

import (
	"net/http"
	"notes/model"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Summary Get Item list
// @Security ApiKeyAuth
// @Tags items
// @Description get Item list
// @Id get-item
// @Accepr json
// @Produce json
// @Param id path string true "List id"
// @Success 200 {object} model.NoteItem
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/lists/{id}/items [get]
func (h *Handler) getAllItems(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	listId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid list id param")
		return
	}

	items, err := h.usecases.NoteItem.GetAll(userId, listId)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	if len(items) == 0 {
		c.JSON(http.StatusOK, statusResponse{
			Status: "not item in list",
		})
		return
	}

	c.JSON(http.StatusOK, items)
}

// @Summary Create Item list
// @Security ApiKeyAuth
// @Tags items
// @Description create Item list
// @Id create-item
// @Accepr json
// @Produce json
// @Param id path string true "List id"
// @Param input body model.ItemInput true "list info"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/lists/{id}/items [post]
func (h *Handler) createItem(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	listId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid list id param")
		return
	}

	var item model.NoteItem
	if err := c.BindJSON(&item); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	itemId, err := h.usecases.NoteItem.Create(userId, listId, item)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": itemId,
	})
}

// @Summary Get Item By Id
// @Security ApiKeyAuth
// @Tags items
// @Description get item by id
// @ID get-item-by-id
// @Accept json
// @Produce json
// @Param id path string true "Item id"
// @Success 200 {object} model.NoteItem
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/items/{id} [get]
func (h *Handler) getItemById(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	itemId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid item id param")
		return
	}

	item, err := h.usecases.NoteItem.GetItemById(userId, itemId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, item)

}

// @Summary Update Item by id
// @Security ApiKeyAuth
// @Tags items
// @Description update Item by id
// @Id update-item-by-id
// @Accepr json
// @Produce json
// @Param id path string true "Item id"
// @Param input body model.ItemInput true "list info"
// @Success 200 {object} statusResponse
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/items/{id} [put]
func (h *Handler) updateItem(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	itemId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid item id param")
		return
	}

	var item model.ItemInput
	if err := c.BindJSON(&item); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.usecases.NoteItem.Update(userId, itemId, item); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}

// @Summary Delete Item By Id
// @Security ApiKeyAuth
// @Tags items
// @Description delete items by id
// @ID delete-items-by-id
// @Accept json
// @Produce json
// @Param id path string true "Item id"
// @Success 200 {object} statusResponse
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/items/{id} [delete]
func (h *Handler) deleteItem(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	itemId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid item id param")
		return
	}

	if err := h.usecases.NoteItem.Delete(userId, itemId); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}
