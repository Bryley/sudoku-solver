/*
Copyright © 2023 Bryley Manning-Hayter bryleyhayter@gmail.com
*/
package cmd

import (
	// "fmt"

	"github.com/Bryley/sudoku-solver/objects"
	"github.com/Bryley/sudoku-solver/solver"
	"github.com/spf13/cobra"
)

// solveCmd represents the solve command
var solveCmd = &cobra.Command{
	Use:   "solve",
    Args: cobra.ExactArgs(1),
	Short: "Solves a sudoku from a filename",
	Long: `Solves a sudoku from a file, The format of the file must contain the
numbers in order reading from top left and will read '#' symbols as empty
and numbers as a number for that cell, every other character including
newlines are ignored`,
	Run: func(cmd *cobra.Command, args []string) {
        board, _ := objects.NewBoardFromFile(args[0])
        solver.SolveBoard(board)
        board.PrintBoard()
	},
}

func init() {
	rootCmd.AddCommand(solveCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// solveCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// solveCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
