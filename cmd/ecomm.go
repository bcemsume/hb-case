package main

import (
	"fmt"
	"hbcase/aggregate"
	"hbcase/console"
	"hbcase/services"
	"strconv"

	"github.com/spf13/cobra"
)

func main() {
	var createProduct = &cobra.Command{
		Use:   "create_product [create new product]",
		Short: "Usage: create_product P1 100 1000",
		Args:  cobra.MinimumNArgs(3),
		Run: func(cmd *cobra.Command, args []string) {
			ecs, e := services.NewECommService(services.WithSimdbProductRepository())
			if e != nil {
				console.WriteError(e.Error())
			}
			code := args[0]
			quantity, _ := strconv.Atoi(args[1])
			price, _ := strconv.ParseFloat(args[2], 64)
			prd, e := aggregate.NewProduct(code, quantity, price)
			if e != nil {
				console.WriteError(e.Error())
			}
			err := ecs.Products.Save(prd)
			if err != nil {
				console.WriteError(err.Error())
			}
			console.WriteInfo(fmt.Sprintf("Product created; Code %s, price %f, quantity %d", code, price, quantity))
		},
	}

	var createOrder = &cobra.Command{
		Use:   "create_order [create_order P1 3]",
		Short: "Usage: create_order P1 3",
		Args:  cobra.MinimumNArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			ecs, e := services.NewECommService(services.WithSimdbOrderRepository(),
				services.WithSimdbProductRepository(),
			)
			if e != nil {
				console.WriteError(e.Error())
			}
			code := args[0]
			quantity, _ := strconv.Atoi(args[1])

			err := ecs.CreateOrder(code, quantity)

			if err != nil {
				console.WriteError(err.Error())
			}
			console.WriteInfo(fmt.Sprintf("Order created; Code %s, quantity %d", code, quantity))

		},
	}

	var createCampaign = &cobra.Command{
		Use:   "create_campaign [create_campaign C1 P1 10 20 100]",
		Short: "Usage: create_campaign C1 P1 10 20 100",
		Args:  cobra.MinimumNArgs(5),
		Run: func(cmd *cobra.Command, args []string) {
			ecs, e := services.NewECommService(services.WithSimdbCampaignRepository(), services.WithSimdbProductRepository())
			if e != nil {
				console.WriteError(e.Error())
			}
			name := args[0]
			productCode := args[1]
			duration, _ := strconv.Atoi(args[2])
			limit, _ := strconv.Atoi(args[2])
			count, _ := strconv.Atoi(args[2])

			err := ecs.CreateCampaign(name, productCode, duration, limit, count)
			if err != nil {
				console.WriteError(err.Error())
				return
			}

			console.WriteInfo(fmt.Sprintf("Campaign created; name %s, product %s, duration %d, limit %d, target sales count %d", name, productCode, duration, limit, count))

		},
	}

	var getProductInfo = &cobra.Command{
		Use:   "get_product_info  [get_product_info  P1]",
		Short: "Usage: get_product_info P1",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			ecs, e := services.NewECommService(services.WithSimdbProductRepository(),
				services.WithSimdbCampaignRepository(),
				services.WithSimdbTimeRepository(),
				services.WithSimdbOrderRepository())
			if e != nil {
				console.WriteError(e.Error())
			}
			code := args[0]
			p, e := ecs.GetProductInfo(code)
			if e != nil {
				console.WriteError(e.Error())
				return
			}
			console.WriteInfo(p)

		},
	}

	var getCampaignInfo = &cobra.Command{
		Use:   "get_campaign_info [get_campaign_info C1]",
		Short: "Usage: get_campaign_info C1",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			ecs, e := services.NewECommService(services.WithSimdbCampaignRepository())
			if e != nil {
				console.WriteError(e.Error())
			}
			code := args[0]

			println(ecs.GetCampaignInfo(code))

		},
	}

	var increaseTime = &cobra.Command{
		Use:   "increase_time [increase_time 1]",
		Short: "Usage: increase_time 1",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			ecs, e := services.NewECommService(services.WithSimdbCampaignRepository(), services.WithSimdbTimeRepository())
			time, _ := strconv.Atoi(args[0])

			if e != nil {
				console.WriteError(e.Error())
			}
			t := aggregate.NewTime(time)

			tim, err := ecs.Time.IncreaseTime(t)

			if err != nil {
				console.WriteError(err.Error())
				return
			}

			console.WriteInfo(fmt.Sprintf("Time increased; time is 0%d:00", tim.GetTime()))

		},
	}

	var rootCmd = &cobra.Command{Use: "hbEComm"}
	rootCmd.AddCommand(createProduct, createOrder, createCampaign, getProductInfo, getCampaignInfo, increaseTime)
	rootCmd.Execute()
}
