package main

import (
	"fmt"

	"github.com/gen2brain/raylib-go/raylib"
)

type state struct {
	trx                *trx
	rec                *rec
	truckSpeed         float32
	currentTruckFrame  uint8
	sighnX             float32
	isGameOver         bool
	isDuckDucking      bool
	score              uint32
	hasSighnPassDuck   bool
	planeX             float32
	hasPlanePassDuck   bool
	planeSpawnlocation float32
	sighnSpawnLocation float32
	duckY              float32
	isDuckJumping      bool
	duckJumpTime       float32
	duckRotation       float32
	difficultyTimer    float32
	sighnSpeed         float32
	planeSpeed         float32
}

type trx struct {
	backGround          rl.Texture2D
	truckTexture        rl.Texture2D
	duckTextureStanding rl.Texture2D
	duckDucking         rl.Texture2D
	roadSighnTrx        rl.Texture2D
	plane               rl.Texture2D
}

type rec struct {
	truckRec     rl.Rectangle
	duckRec      rl.Rectangle
	duckDucking  rl.Rectangle
	roadSighnRec rl.Rectangle
	plane        rl.Rectangle
}

func main() {
	rl.InitWindow(800, 450, "!DuckOnATruck!")
	rl.SetTargetFPS(60)

	trx := &trx{
		backGround:          rl.LoadTexture("Resources/map.png"),
		truckTexture:        rl.LoadTexture("Resources/Truck.png"),
		duckTextureStanding: rl.LoadTexture("Resources/DuckStanding.png"),
		duckDucking:         rl.LoadTexture("Resources/duckDucking.png"),
		roadSighnTrx:        rl.LoadTexture("Resources/RoadSighn.png"),
		plane:               rl.LoadTexture("Resources/Plane.png"),
	}

	rec := &rec{
		truckRec:     recMaker(0, float32(trx.truckTexture.Width)/2, 160),
		duckRec:      recMaker(0, 80, 80),
		duckDucking:  recMaker(0, 80, 80),
		roadSighnRec: recMaker(0, float32(trx.roadSighnTrx.Width), float32(trx.roadSighnTrx.Height)),
		plane:        recMaker(0, float32(trx.plane.Width), float32(trx.plane.Height)),
	}

	state := &state{
		trx:                trx,
		rec:                rec,
		truckSpeed:         0,
		currentTruckFrame:  0,
		sighnX:             40,
		isGameOver:         false,
		isDuckDucking:      false,
		score:              0,
		hasSighnPassDuck:   true,
		hasPlanePassDuck:   false,
		planeX:             700,
		planeSpawnlocation: float32(rl.GetRandomValue(900, 2200)),
		sighnSpawnLocation: float32(rl.GetRandomValue(900, 2200)),
		duckY:              240,
		isDuckJumping:      false,
		duckJumpTime:       0,
		duckRotation:       0,
		difficultyTimer:    0,
		sighnSpeed:         100,
		planeSpeed:         400,
	}

	for !rl.WindowShouldClose() {
		update(state)

		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)

		rl.DrawTexture(trx.backGround, 0, 0, rl.White)
		rl.DrawTextureRec(trx.truckTexture, rec.truckRec, rl.NewVector2(300, 250), rl.White)
		if state.isDuckDucking {
			rl.DrawTextureRec(trx.duckDucking, rec.duckDucking, rl.NewVector2(320, state.duckY), rl.White)
		} else {
			rl.DrawTextureRec(trx.duckTextureStanding, rec.duckRec, rl.NewVector2(320, state.duckY-5), rl.White)
		}
		rl.DrawTextureRec(trx.roadSighnTrx, rec.roadSighnRec, rl.NewVector2(state.sighnX, 130), rl.White)
		rl.DrawTextureRec(trx.plane, rec.plane, rl.NewVector2(state.planeX, 220), rl.White)
		rl.DrawText(fmt.Sprintf("Score: %d", state.score), 340, 0, 30, rl.Black)

		if state.isGameOver {
			rl.DrawRectangle(0, 0, 850, 450, rl.White)
			rl.DrawText("    Game Is Over\nPress R To Restart", 250, 150, 32, rl.Red)
			rl.DrawTexturePro(trx.duckTextureStanding, rec.duckRec, rl.NewRectangle((850/2)-(rec.duckRec.Width/2), 325, 80, 80), rl.NewVector2(40, 40), state.duckRotation, rl.White)
		}

		rl.EndDrawing()
	}
	rl.UnloadTexture(trx.plane)
	rl.UnloadTexture(trx.roadSighnTrx)
	rl.UnloadTexture(trx.duckTextureStanding)
	rl.UnloadTexture(trx.truckTexture)
	rl.UnloadTexture(trx.backGround)
	rl.CloseWindow()
}

func recMaker(x float32, Width float32, Height float32) rl.Rectangle {
	return rl.Rectangle{
		X:      x,
		Y:      0,
		Width:  Width,
		Height: Height,
	}
}

func restart(state *state) {
	state.rec.duckRec.Y = 80
	state.rec.duckDucking.Y = 80
	state.sighnX = state.sighnSpawnLocation
	state.score = 0
	state.planeX = state.planeSpawnlocation
	state.hasSighnPassDuck = false
	state.hasPlanePassDuck = false
	state.duckRotation = 0
	state.difficultyTimer = 0
	state.planeSpeed = 400
	state.sighnSpeed = 100
}

func update(state *state) {

	var dt = rl.GetFrameTime()

	state.truckSpeed += dt

	if state.truckSpeed >= 0.3 {
		state.truckSpeed = 0
		if state.currentTruckFrame == 1 {
			state.currentTruckFrame = 0
		} else {
			state.currentTruckFrame = 1
		}
		state.rec.truckRec.X = (state.rec.truckRec.Width * float32(state.currentTruckFrame))
	}

	if rl.IsKeyPressed(rl.KeyR) {
		state.isGameOver = false
		restart(state)
	}
	if state.isGameOver {
		state.duckRotation += 100 * dt
		if state.duckRotation > 360 {
			state.duckRotation = 0
		}
		return
	}
	if rl.IsKeyDown(rl.KeyK) {
		state.isDuckDucking = true
	} else {
		state.isDuckDucking = false
	}

	if state.sighnX <= 0-state.rec.roadSighnRec.Width {
		state.sighnX = state.sighnSpawnLocation
		state.hasSighnPassDuck = false
	}
	if state.planeX <= 0-state.rec.plane.Width {
		state.planeX = state.planeSpawnlocation
		state.hasPlanePassDuck = false
	}

	if !state.hasSighnPassDuck && state.sighnX <= 425-80-state.rec.roadSighnRec.Width {
		state.hasSighnPassDuck = true
		state.score += 1
	}
	if !state.hasPlanePassDuck && state.planeX <= 425-80-state.rec.plane.Width {
		state.hasPlanePassDuck = true
		state.score += 1
	}

	if !state.isDuckJumping {
		if rl.IsKeyDown(rl.KeyJ) {
			state.isDuckJumping = true
			state.duckY -= 60
			state.duckJumpTime = 0.7
		}
	} else {
		state.duckJumpTime -= dt
		if state.duckJumpTime <= 0 || rl.IsKeyPressed(rl.KeyK) {
			state.duckJumpTime = 0
			state.isDuckJumping = false
			state.duckY += 60
		}
	}

	if !state.isDuckDucking && rl.CheckCollisionRecs(rl.Rectangle{
		X:      320 + 36,
		Y:      state.duckY,
		Width:  41,
		Height: state.rec.duckRec.Height,
	}, rl.Rectangle{
		X:      state.sighnX + 24,
		Y:      130,
		Width:  32,
		Height: 128,
	}) {
		state.isGameOver = true
	}

	if !state.isDuckJumping && rl.CheckCollisionRecs(rl.Rectangle{
		X:      320 + 36,
		Y:      state.duckY,
		Width:  41,
		Height: state.rec.duckRec.Height,
	}, rl.Rectangle{
		X:      state.planeX + 10,
		Y:      220 + 36,
		Width:  84,
		Height: 52,
	}) {
		state.isGameOver = true
	}

	state.difficultyTimer += dt

	if state.difficultyTimer > 1 {
		state.difficultyTimer = 0
		state.sighnSpeed += 25
		state.planeSpeed += 25
	}

	state.sighnX -= state.sighnSpeed * dt
	state.planeX -= state.planeSpeed * dt

}
