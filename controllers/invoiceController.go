package controllers

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/simpleshaik1/restuarant-management/database"
	"github.com/simpleshaik1/restuarant-management/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/http"
	"time"
)

type InvoiceViewFormat struct {
	Invoice_id       string
	Payment_method   string
	Order_id         string
	Payment_status   *string
	Payment_due      interface{}
	Table_number     interface{}
	Payment_due_date time.Time
	Oder_details     interface{}
}

var InvoiceCollection *mongo.Collection = database.OpenCollection(database.Client, "invoice")

func GetInvoices() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		result, err := InvoiceCollection.Find(ctx, bson.D{})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error occured while listing invoice items"})
		}
		var allInvoices []bson.M
		if err = result.All(ctx, &allInvoices); err != nil {
			log.Fatal(err)
		}
		c.JSON(http.StatusOK, allInvoices)
	}
}

func GetInvoice() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		invoiceId := c.Param("invoice_id")

		var invoice models.Invoice

		err := InvoiceCollection.FindOne(ctx, bson.M{"invoice_id": invoiceId}).Decode(&invoice)
		defer cancel()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error occurred while getting invoice"})
		}

		var invoiceView InvoiceViewFormat

		allOrderItems, err := ItemsByOrder(invoice.Order_id)
		invoiceView.Order_id = invoice.Order_id
		invoiceView.Payment_due_date = invoice.Payment_due_date

		invoiceView.Payment_method = "null"
		if invoice.Payment_method != nil {
			invoiceView.Payment_method = *invoice.Payment_method
		}

		invoiceView.Invoice_id = invoice.Invoice_id
		invoiceView.Payment_status = *&invoice.Payment_status
		invoiceView.Payment_due = allOrderItems[0]["payment_due"]
		invoiceView.Table_number = allOrderItems[0]["table_number"]
		invoiceView.Oder_details = allOrderItems[0]["oder_details"]

		c.JSON(http.StatusOK, invoiceView)
	}
}

func CreateInvoice() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		var invoice models.Invoice
		if err := c.BindJSON(&invoice); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		var order models.Order

		err := orderCollection.FindOne(ctx, bson.M{"order_id": invoice.Order_id}).Decode(&order)
		if err != nil {
			msg := fmt.Sprintf("Order was not found")
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}
	}
}

func UpdateInvoice() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		var invoice models.Invoice
		invoiceId := c.Param("invoice_id")

		if err := c.ShouldBindJSON(&invoice); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		filter := bson.M{"invoice_id": invoiceId}

		var updateObj primitive.D

		if invoice.Payment_method != nil {
		}
		if invoice.Payment_status != nil {
		}

		invoice.CreatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		updateObj = append(updateObj, bson.E{"updated_at", invoice.UpdatedAt})

		upsert := true

		opt := options.UpdateOptions{
			Upsert: &upsert,
		}
		status := "PENDING"
		if invoice.Payment_status == nil {
			invoice.Payment_status = &status
		}
		result, err := InvoiceCollection.UpdateOne(ctx, filter, bson.M{"$set": updateObj}, &opt)

		if err != nil {
			msg := fmt.Sprintf("Error occured while updating invoice")
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}
		c.JSON(http.StatusOK, result)
	}
}
