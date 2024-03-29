package main

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func insertGame(game *Game, db *sql.DB) error {
	// TODO: Prepare just once
	stmt := Statements["insertGame"]

	_, err := stmt.Exec(
		game.Date,
		game.NumberOfGame,
		game.DayOfWeek,
		game.VisitingTeam,
		game.VisitingTeamLeague,
		game.VisitingGameNumber,
		game.HomeTeam,
		game.HomeTeamLeague,
		game.HomeTeamGameNumber,
		game.VisitingTeamScore,
		game.HomeTeamScore,
		game.GameLengthInOuts,
		game.DayNightIndicator,
		game.CompletionInformation,
		game.ForfeitInformation,
		game.ProtestInformation,
		game.ParkID,
		game.Attendance,
		game.TimeOfGameInMins,
		game.VisitingLineScore,
		game.HomeLineScore,
		game.VisitingAB,
		game.VisitingH,
		game.Visiting2B,
		game.Visiting3B,
		game.VisitingHR,
		game.VisitingRBI,
		game.VisitingSH,
		game.VisitingSF,
		game.VisitingHBP,
		game.VisitingBB,
		game.VisitingIBB,
		game.VisitingK,
		game.VisitingSB,
		game.VisitingCS,
		game.VisitingGIDP,
		game.VisitingCI,
		game.VisitingLOB,
		game.VisitingPitchersUsed,
		game.VisitingIndividualEarnedRuns,
		game.VisitingTeamEarnedRuns,
		game.VisitingWildPitches,
		game.VisitingBalks,
		game.VisitingPutouts,
		game.VisitingAssists,
		game.VisitingErrors,
		game.VisitingPassedBalls,
		game.VisitingDoublePlays,
		game.VisitingTriplePlays,
		game.HomeAB,
		game.HomeH,
		game.Home2B,
		game.Home3B,
		game.HomeHR,
		game.HomeRBI,
		game.HomeSH,
		game.HomeSF,
		game.HomeHBP,
		game.HomeBB,
		game.HomeIBB,
		game.HomeK,
		game.HomeSB,
		game.HomeCS,
		game.HomeGIDP,
		game.HomeCI,
		game.HomeLOB,
		game.HomePitchersUsed,
		game.HomeIndividualEarnedRuns,
		game.HomeTeamEarnedRuns,
		game.HomeWildPitches,
		game.HomeBalks,
		game.HomePutouts,
		game.HomeAssists,
		game.HomeErrors,
		game.HomePassedBalls,
		game.HomeDoublePlays,
		game.HomeTriplePlays,
		game.HomePlateUmpireID,
		game.HomePlateUmpireName,
		game.FirstBaseUmpireID,
		game.FirstBaseUmpireName,
		game.SecondBaseUmpireID,
		game.SecondBaseUmpireName,
		game.ThirdBaseUmpireID,
		game.ThirdBaseUmpireName,
		game.LeftFieldUmpireID,
		game.LeftFieldUmpireName,
		game.RightFieldUmpireID,
		game.RightFieldUmpireName,
		game.VisitingManagerID,
		game.VisitingManagerName,
		game.HomeManagerID,
		game.HomeManagerName,
		game.WinningPitcherID,
		game.WinningPitcherName,
		game.LosingPitcherID,
		game.LosingPitcherName,
		game.SavingPitcherID,
		game.SavingPitcherName,
		game.GameWinningRBIBatterID,
		game.GameWinningRBIBatterName,
		game.VisitingStartingPitcherID,
		game.VisitingStartingPitcherName,
		game.HomeStartingPitcherID,
		game.HomeStartingPitcherName,
		game.VisitingPlayer1ID,
		game.VisitingPlayer1Name,
		game.VisitingPlayer1Position,
		game.VisitingPlayer2ID,
		game.VisitingPlayer2Name,
		game.VisitingPlayer2Position,
		game.VisitingPlayer3ID,
		game.VisitingPlayer3Name,
		game.VisitingPlayer3Position,
		game.VisitingPlayer4ID,
		game.VisitingPlayer4Name,
		game.VisitingPlayer4Position,
		game.VisitingPlayer5ID,
		game.VisitingPlayer5Name,
		game.VisitingPlayer5Position,
		game.VisitingPlayer6ID,
		game.VisitingPlayer6Name,
		game.VisitingPlayer6Position,
		game.VisitingPlayer7ID,
		game.VisitingPlayer7Name,
		game.VisitingPlayer7Position,
		game.VisitingPlayer8ID,
		game.VisitingPlayer8Name,
		game.VisitingPlayer8Position,
		game.VisitingPlayer9ID,
		game.VisitingPlayer9Name,
		game.VisitingPlayer9Position,
		game.HomePlayer1ID,
		game.HomePlayer1Name,
		game.HomePlayer1Position,
		game.HomePlayer2ID,
		game.HomePlayer2Name,
		game.HomePlayer2Position,
		game.HomePlayer3ID,
		game.HomePlayer3Name,
		game.HomePlayer3Position,
		game.HomePlayer4ID,
		game.HomePlayer4Name,
		game.HomePlayer4Position,
		game.HomePlayer5ID,
		game.HomePlayer5Name,
		game.HomePlayer5Position,
		game.HomePlayer6ID,
		game.HomePlayer6Name,
		game.HomePlayer6Position,
		game.HomePlayer7ID,
		game.HomePlayer7Name,
		game.HomePlayer7Position,
		game.HomePlayer8ID,
		game.HomePlayer8Name,
		game.HomePlayer8Position,
		game.HomePlayer9ID,
		game.HomePlayer9Name,
		game.HomePlayer9Position,
		game.AdditionalInformation,
		game.AcquisitionInformation,
	)

	return err
}
