/*
Copyright © 2023 Bryley Manning-Hayter bryleyhayter@gmail.com
*/
package cmd

import (
	"fmt"

	"github.com/Bryley/sudoku-solver/objects"
	"github.com/Bryley/sudoku-solver/solver"
	"github.com/spf13/cobra"
)

// isSolvedCmd represents the isSolved command
var isSolvedCmd = &cobra.Command{
	Use:   "is-solved",
    Args: cobra.ExactArgs(1),
	Short: "Checks if a suduko is solved or not",
	Long: `Will calculate the numbers used to see if the suduko is already solved
or not`,
	Run: func(cmd *cobra.Command, args []string) {
        board, _ := objects.NewBoardFromFile(args[0])
        fmt.Print("Board is ")
        if !solver.IsSolved(board) {
            fmt.Print("not ")
        }
        fmt.Println("solved")
	},
}

func init() {
	rootCmd.AddCommand(isSolvedCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// isSolvedCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// isSolvedCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
