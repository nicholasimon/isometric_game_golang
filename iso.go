package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var ( // MARK: var

	// timers
	framecount int
	// player

	nextblocknumber           int
	pblocknext                int
	pmoveon, pmove2on, checkl bool
	direction, path           string
	pblock                    = 15
	pblocknew                 = pblock
	// images
	tree1 = rl.NewRectangle(0, 102, 47, 73)
	ter1  = rl.NewRectangle(0, 0, 100, 50)
	ter2  = rl.NewRectangle(0, 50, 100, 50)
	imgs  rl.Texture2D
	// map
	blocks      = make([]bool, 450)
	blockactive int
	levelmap    = make([]string, 450)
	lineswitch  bool

	// input
	mousepos rl.Vector2
	// debug
	debugon bool
	// camera screen
	borderson, gridon bool
	screenW           = int32(1600)
	screenH           = int32(900)
	camera            rl.Camera2D
)

// MARK: notes
/*
	TOTAL 450 blocks
	OUTER	15 horiz X 16 vert = 240 blocks / 100px W X 50px H
	INNER 	14 horiz X 15 vert = 210 blocks / 100px W X 50px H

*/
func start() { // MARK: start
	camera.Zoom = 1.0
	debugon = true
	//borderson = true
	gridon = true
}
func input() { // MARK: keys input

	if rl.IsKeyDown(rl.KeyUp) {
		camera.Target.Y -= 10
	}
	if rl.IsKeyDown(rl.KeyDown) {
		camera.Target.Y += 10
	}
	if rl.IsKeyDown(rl.KeyRight) {
		camera.Target.X += 10
	}
	if rl.IsKeyDown(rl.KeyLeft) {
		camera.Target.X -= 10
	}
	if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
		pblocknew = blockactive
		if pblocknew != pblock {
			choosedirection()
		}
	}

	if rl.IsKeyPressed(rl.KeyKpMultiply) {
		if gridon {
			gridon = false
		} else {
			gridon = true
		}
	}
	if rl.IsKeyPressed(rl.KeyKpDivide) {
		if borderson {
			borderson = false
		} else {
			borderson = true
		}
	}

	if rl.IsKeyPressed(rl.KeyKpAdd) {
		if camera.Zoom < 1.8 {
			if camera.Zoom == 1.0 {
				camera.Zoom = 1.3
			} else if camera.Zoom == 1.3 {
				camera.Zoom = 1.8
			}
		}
	}
	if rl.IsKeyPressed(rl.KeyKpSubtract) {
		if camera.Zoom > 1.0 {
			if camera.Zoom == 1.8 {
				camera.Zoom = 1.3
			} else if camera.Zoom == 1.3 {
				camera.Zoom = 1.0
			}
		}
	}

	if rl.IsKeyPressed(rl.KeyKpDecimal) {
		if debugon {
			debugon = false
		} else {
			debugon = true
		}
	}

}
func choosedirection() { // MARK: choosedirection()

	if pblocknew != pblock {

		if pblocknew < pblock { // above/behind player

			if pblock-pblocknew < 12 {
				direction = "l"
				pmoveon = true
			} else if (pblock-pblocknew)%29 == 0 {
				direction = "u"
				pmoveon = true
			} else if (pblock-pblocknew)%14 == 0 {
				direction = "ur"
				pmoveon = true
			} else if (pblock-pblocknew)%15 == 0 {
				direction = "ul"
				pmoveon = true
			} else {
				findpath()
			}

		} else { // below/in front player
			if pblocknew-pblock < 12 {
				direction = "r"
				pmoveon = true
			} else if (pblocknew-pblock)%29 == 0 {
				direction = "d"
				pmoveon = true
			} else if (pblocknew-pblock)%15 == 0 {
				direction = "dr"
				pmoveon = true
			} else if (pblocknew-pblock)%14 == 0 {
				direction = "dl"
				pmoveon = true
			} else {
				findpath()
			}
		}
	}
}
func findpath() { // MARK: findpath()

	if pblocknew < pblock { // above/behind player

		checkl = false

		check1 := pblock - 15
		check2 := pblock - 15*2
		check3 := pblock - 15*3
		check4 := pblock - 15*4
		check5 := pblock - 15*5
		check6 := pblock - 15*6
		check7 := pblock - 15*7
		check8 := pblock - 15*8
		check9 := pblock - 15*9
		check10 := pblock - 15*10

		if check1 > 0 {
			if (check1-pblocknew)%14 == 0 {
				pblocknext = check1
				pmove2on = true
				direction = "ul"
				checkl = true
			}
		}
		if check2 > 0 {
			if (check2-pblocknew)%14 == 0 {
				pblocknext = check2
				pmove2on = true
				direction = "ul"
				checkl = true
			}
		}
		if check3 > 0 {
			if (check3-pblocknew)%14 == 0 {
				pblocknext = check3
				pmove2on = true
				direction = "ul"
				checkl = true
			}
		}
		if check4 > 0 {
			if (check4-pblocknew)%14 == 0 {
				pblocknext = check4
				pmove2on = true
				direction = "ul"
				checkl = true
			}
		}
		if check5 > 0 {
			if (check5-pblocknew)%14 == 0 {
				pblocknext = check5
				pmove2on = true
				direction = "ul"
				checkl = true
			}
		}
		if check6 > 0 {
			if (check6-pblocknew)%14 == 0 {
				pblocknext = check6
				pmove2on = true
				direction = "ul"
				checkl = true
			}
		}
		if check7 > 0 {
			if (check7-pblocknew)%14 == 0 {
				pblocknext = check7
				pmove2on = true
				direction = "ul"
				checkl = true
			}
		}
		if check8 > 0 {
			if (check8-pblocknew)%14 == 0 {
				pblocknext = check8
				pmove2on = true
				direction = "ul"
				checkl = true
			}
		}
		if check9 > 0 {
			if (check9-pblocknew)%14 == 0 {
				pblocknext = check9
				pmove2on = true
				direction = "ul"
				checkl = true
			}
		}
		if check10 > 0 {
			if (check10-pblocknew)%14 == 0 {
				pblocknext = check10
				pmove2on = true
				direction = "ul"
				checkl = true
			}
		}

		if checkl == false {

			check1 := pblock - 14
			check2 := pblock - 14*2
			check3 := pblock - 14*3
			check4 := pblock - 14*4
			check5 := pblock - 14*5
			check6 := pblock - 14*6
			check7 := pblock - 14*7
			check8 := pblock - 14*8
			check9 := pblock - 14*9
			check10 := pblock - 14*10

			if check1 > 0 {
				if (check1-pblocknew)%15 == 0 {
					pblocknext = check1
					pmove2on = true
					direction = "ur"

				}
			}
			if check2 > 0 {
				if (check2-pblocknew)%15 == 0 {
					pblocknext = check2
					pmove2on = true
					direction = "ur"

				}
			}
			if check3 > 0 {
				if (check3-pblocknew)%15 == 0 {
					pblocknext = check3
					pmove2on = true
					direction = "ur"

				}
			}
			if check4 > 0 {
				if (check4-pblocknew)%15 == 0 {
					pblocknext = check4
					pmove2on = true
					direction = "ur"

				}
			}
			if check5 > 0 {
				if (check5-pblocknew)%15 == 0 {
					pblocknext = check5
					pmove2on = true
					direction = "ur"

				}
			}
			if check6 > 0 {
				if (check6-pblocknew)%15 == 0 {
					pblocknext = check6
					pmove2on = true
					direction = "ur"

				}
			}
			if check7 > 0 {
				if (check7-pblocknew)%15 == 0 {
					pblocknext = check7
					pmove2on = true
					direction = "ur"

				}
			}
			if check8 > 0 {
				if (check8-pblocknew)%15 == 0 {
					pblocknext = check8
					pmove2on = true
					direction = "ur"

				}
			}
			if check9 > 0 {
				if (check9-pblocknew)%15 == 0 {
					pblocknext = check9
					pmove2on = true
					direction = "ur"

				}
			}
			if check10 > 0 {
				if (check10-pblocknew)%15 == 0 {
					pblocknext = check10
					pmove2on = true
					direction = "ur"

				}
			}

		}

	} else { // below/in front player

		checkl = false

		check1 := pblock + 14
		check2 := pblock + 14*2
		check3 := pblock + 14*3
		check4 := pblock + 14*4
		check5 := pblock + 14*5
		check6 := pblock + 14*6
		check7 := pblock + 14*7
		check8 := pblock + 14*8
		check9 := pblock + 14*9
		check10 := pblock + 14*10

		if check1 < 450 {
			if (pblocknew-check1)%29 == 0 {
				pblocknext = check1
				pmove2on = true
				direction = "dl"
				checkl = true
			}
		}
		if check2 < 450 {
			if (pblocknew-check2)%29 == 0 {
				pblocknext = check2
				pmove2on = true
				direction = "dl"
				checkl = true
			}
		}
		if check3 < 450 {
			if (pblocknew-check3)%29 == 0 {
				pblocknext = check3
				pmove2on = true
				direction = "dl"
				checkl = true
			}
		}
		if check4 < 450 {
			if (pblocknew-check4)%29 == 0 {
				pblocknext = check4
				pmove2on = true
				direction = "dl"
				checkl = true
			}
		}
		if check5 < 450 {
			if (pblocknew-check5)%29 == 0 {
				pblocknext = check5
				pmove2on = true
				direction = "dl"
				checkl = true
			}
		}
		if check6 < 450 {
			if (pblocknew-check6)%29 == 0 {
				pblocknext = check6
				pmove2on = true
				direction = "dl"
				checkl = true
			}
		}
		if check7 < 450 {
			if (pblocknew-check7)%29 == 0 {
				pblocknext = check7
				pmove2on = true
				direction = "dl"
				checkl = true
			}
		}
		if check8 < 450 {
			if (pblocknew-check8)%29 == 0 {
				pblocknext = check8
				pmove2on = true
				direction = "dl"
				checkl = true
			}
		}
		if check9 < 450 {
			if (pblocknew-check9)%29 == 0 {
				pblocknext = check9
				pmove2on = true
				direction = "dl"
				checkl = true
			}
		}
		if check10 < 450 {
			if (pblocknew-check10)%29 == 0 {
				pblocknext = check10
				pmove2on = true
				direction = "dl"
				checkl = true
			}
		}

		if checkl == false {

			check1 := pblock + 15
			check2 := pblock + 15*2
			check3 := pblock + 15*3
			check4 := pblock + 15*4
			check5 := pblock + 15*5
			check6 := pblock + 15*6
			check7 := pblock + 15*7
			check8 := pblock + 15*8
			check9 := pblock + 15*9
			check10 := pblock + 15*10

			if check1 < 450 {
				if (pblocknew-check1)%29 == 0 {
					pblocknext = check1
					pmove2on = true
					direction = "dr"
				}
			}
			if check2 < 450 {
				if (pblocknew-check2)%29 == 0 {
					pblocknext = check2
					pmove2on = true
					direction = "dr"

				}
			}
			if check3 < 450 {
				if (pblocknew-check3)%29 == 0 {
					pblocknext = check3
					pmove2on = true
					direction = "dr"

				}
			}
			if check4 < 450 {
				if (pblocknew-check4)%29 == 0 {
					pblocknext = check4
					pmove2on = true
					direction = "dr"

				}
			}
			if check5 < 450 {
				if (pblocknew-check5)%29 == 0 {
					pblocknext = check5
					pmove2on = true
					direction = "dr"

				}
			}
			if check6 < 450 {
				if (pblocknew-check6)%29 == 0 {
					pblocknext = check6
					pmove2on = true
					direction = "dr"

				}
			}
			if check7 < 450 {
				if (pblocknew-check7)%29 == 0 {
					pblocknext = check7
					pmove2on = true
					direction = "dr"

				}
			}
			if check8 < 450 {
				if (pblocknew-check8)%29 == 0 {
					pblocknext = check8
					pmove2on = true
					direction = "dr"

				}
			}
			if check9 < 450 {
				if (pblocknew-check9)%29 == 0 {
					pblocknext = check9
					pmove2on = true
					direction = "dr"

				}
			}
			if check10 < 450 {
				if (pblocknew-check10)%29 == 0 {
					pblocknext = check10
					pmove2on = true
					direction = "dr"

				}
			}

		}

	}

}
func moveplayer2() { // MARK: moveplayer2()

	switch direction {
	case "ur":
		if pblocknext != pblock {
			if framecount%15 == 0 {
				pblock -= 14
			}
		} else {
			choosedirection()
			pmove2on = false
		}
	case "ul":
		if pblocknext != pblock {
			if framecount%15 == 0 {
				pblock -= 15
			}
		} else {
			choosedirection()
			pmove2on = false
		}
	case "dr":
		if pblocknext != pblock {
			if framecount%15 == 0 {
				pblock += 15
			}
		} else {
			choosedirection()
			pmove2on = false
		}
	case "dl":
		if pblocknext != pblock {
			if framecount%15 == 0 {
				pblock += 14
			}
		} else {
			choosedirection()
			pmove2on = false
		}
	}
}
func moveplayer() { // MARK: moveplayer()

	switch direction {
	case "ur":
		if pblocknew != pblock {
			if framecount%15 == 0 {
				pblock -= 14
			}
		} else {
			pmoveon = false
		}
	case "ul":
		if pblocknew != pblock {
			if framecount%15 == 0 {
				pblock -= 15
			}
		} else {
			pmoveon = false
		}
	case "dr":
		if pblocknew != pblock {
			if framecount%15 == 0 {
				pblock += 15
			}
		} else {
			pmoveon = false
		}
	case "dl":
		if pblocknew != pblock {
			if framecount%15 == 0 {
				pblock += 14
			}
		} else {
			pmoveon = false
		}
	case "u":
		if pblocknew != pblock {
			if framecount%15 == 0 {
				pblock -= 29
			}
		} else {
			pmoveon = false
		}
	case "d":
		if pblocknew != pblock {
			if framecount%15 == 0 {
				pblock += 29
			}
		} else {
			pmoveon = false
		}
	case "r":
		if pblocknew != pblock {
			if framecount%15 == 0 {
				pblock++
			}
		} else {
			pmoveon = false
		}
	case "l":
		if pblocknew != pblock {
			if framecount%15 == 0 {
				pblock--
			}
		} else {
			pmoveon = false
		}
	}
}

func createmap() { // MARK: createmap()

	for a := 0; a < 450; a++ {
		levelmap[a] = "."
	}
	for a := 0; a < 450; a++ {
		if rolldice()+rolldice() == 12 {
			levelmap[a] = "#"
		}
	}

}
func drawconsole() { // MARK:drawconsole()

	count := 0
	for a := 0; a < 450; a++ {
		b := levelmap[a]
		print(b)
		count++
		if lineswitch {
			if count == 14 {
				count = 0
				println()
				lineswitch = false
			}
		} else {
			if count == 15 {
				count = 0
				println()
				lineswitch = true
			}
		}
	}
}

func main() { // MARK: main()
	rand.Seed(time.Now().UnixNano()) // random numbers
	rl.SetTraceLog(rl.LogError)      // hides INFO window

	start()
	createmap()
	// drawconsole()
	raylib()
}

func raylib() { // MARK: raylib

	rl.InitWindow(screenW, screenH, "iso grid")
	rl.SetExitKey(rl.KeyEnd) // key to end the game and close window

	imgs = rl.LoadTexture("imgs.png") // load images
	rl.SetTargetFPS(60)
	for !rl.WindowShouldClose() { // MARK: WindowShouldClose
		framecount++
		mousepos = rl.GetMousePosition()
		input()
		getblock()

		if pmoveon {
			moveplayer()
		} else if pmove2on {
			moveplayer2()
		}

		rl.BeginDrawing()

		rl.ClearBackground(rl.Black)

		rl.BeginMode2D(camera)

		// MARK: draw map layer 1
		blokx := int32(100)
		bloky := int32(75)
		count := 0
		lineswitch = false
		for a := 0; a < 450; a++ {
			if blockactive == a {
				triv1 := rl.NewVector2(float32(blokx-50), float32(bloky))
				triv2 := rl.NewVector2(float32(blokx), float32(bloky-25))
				triv3 := rl.NewVector2(float32(blokx), float32(bloky+25))
				triv4 := rl.NewVector2(float32(blokx+50), float32(bloky))
				rl.DrawTriangle(triv1, triv3, triv2, rl.DarkGreen)
				rl.DrawTriangle(triv2, triv3, triv4, rl.DarkGreen)
			}
			b := levelmap[a]
			if b == "." {
				// rl.DrawCircle(blokx, bloky, 5, rl.Red)
				tilev := rl.NewVector2(float32(blokx-50), float32(bloky-25))
				if blockactive == a {
					rl.DrawTextureRec(imgs, ter1, tilev, rl.Fade(rl.White, 0.4))
				} else {
					rl.DrawTextureRec(imgs, ter1, tilev, rl.White)
				}

			} else if b == "#" {
				//	rl.DrawCircle(blokx, bloky, 5, rl.Blue)
				tilev := rl.NewVector2(float32(blokx-50), float32(bloky-25))

				if blockactive == a {
					rl.DrawTextureRec(imgs, ter2, tilev, rl.Fade(rl.White, 0.4))

				} else {
					rl.DrawTextureRec(imgs, ter2, tilev, rl.White)

				}
			}

			if a == pblock {
				rl.DrawCircle(blokx, bloky, 10, rl.Maroon)
			}

			count++
			blokx += 100
			if lineswitch {
				if count == 14 {
					count = 0
					blokx = 100
					bloky += 25
					lineswitch = false
				}
			} else {
				if count == 15 {
					count = 0
					blokx = 150
					bloky += 25
					lineswitch = true
				}
			}

		} // draw map layer 1

		// MARK: draw map layer 2
		blokx = int32(100)
		bloky = int32(75)
		count = 0
		lineswitch = false
		for a := 0; a < 450; a++ {

			b := levelmap[a]
			if b == "#" {

				//	treev := rl.NewVector2(float32(blokx), float32(bloky))
				if blockactive == a {

					//	rl.DrawTextureRec(imgs, tree1, treev, rl.Fade(rl.White, 0.4))
				} else {

					//	rl.DrawTextureRec(imgs, tree1, treev, rl.White)
				}
			}

			if a == pblock {
				rl.DrawCircle(blokx, bloky, 10, rl.Maroon)
			}

			count++
			blokx += 100
			if lineswitch {
				if count == 14 {
					count = 0
					blokx = 100
					bloky += 25
					lineswitch = false
				}
			} else {
				if count == 15 {
					count = 0
					blokx = 150
					bloky += 25
					lineswitch = true
				}
			}

		} // draw map layer 2

		// draw grid
		if gridon {
			ychange := int32(0)
			xchange := int32(0)

			for b := 0; b < 15; b++ {
				for a := 0; a < 16; a++ {
					rl.DrawLine(50+xchange, 75+ychange, 100+xchange, 100+ychange, rl.Fade(rl.Black, 0.3))
					rl.DrawLine(100+xchange, 100+ychange, 150+xchange, 75+ychange, rl.Fade(rl.Black, 0.3))
					rl.DrawLine(50+xchange, 75+ychange, 100+xchange, 50+ychange, rl.Fade(rl.Black, 0.3))
					rl.DrawLine(100+xchange, 50+ychange, 150+xchange, 75+ychange, rl.Fade(rl.Black, 0.3))
					ychange += 50
				}
				xchange += 100
				ychange = 0
			}
		} // draw grid

		rl.EndMode2D()

		if borderson {
			rl.DrawRectangle(0, 0, screenW, 75, rl.Fade(rl.Black, 0.9))
			rl.DrawRectangle(0, 75, 100, screenH-75, rl.Fade(rl.Black, 0.9))
			rl.DrawRectangle(screenW-100, 75, 100, screenH-75, rl.Fade(rl.Black, 0.9))
			rl.DrawRectangle(0, screenH-75, screenW, 75, rl.Fade(rl.Black, 0.9))
		}

		if debugon { // MARK: debug
			rl.DrawRectangle(screenW-300, 0, 500, screenW, rl.Fade(rl.Black, 0.8))

			for a := 0; a < 450; a++ {
				if blocks[a] {
					blockactive = a
				}
			}

			blockactiveTEXT := strconv.Itoa(blockactive)
			mouseposXTEXT := fmt.Sprintf("%.0f", mousepos.X)
			mouseposYTEXT := fmt.Sprintf("%.0f", mousepos.Y)
			pblockTEXT := strconv.Itoa(pblock)
			pblocknewTEXT := strconv.Itoa(pblocknew)

			rl.DrawText(mouseposXTEXT, screenW-290, 10, 10, rl.White)
			rl.DrawText("mouseposX", screenW-200, 10, 10, rl.White)
			rl.DrawText(mouseposYTEXT, screenW-290, 20, 10, rl.White)
			rl.DrawText("mouseposY", screenW-200, 20, 10, rl.White)
			rl.DrawText(blockactiveTEXT, screenW-290, 30, 10, rl.White)
			rl.DrawText("blockactive", screenW-200, 30, 10, rl.White)
			rl.DrawText(pblockTEXT, screenW-290, 40, 10, rl.White)
			rl.DrawText("pblock", screenW-200, 40, 10, rl.White)
			rl.DrawText(pblocknewTEXT, screenW-290, 50, 10, rl.White)
			rl.DrawText("pblocknew", screenW-200, 50, 10, rl.White)
			rl.DrawText(direction, screenW-290, 60, 10, rl.White)
			rl.DrawText("direction", screenW-200, 60, 10, rl.White)

			rl.DrawText("grid on/off: * keypad", screenW-290, 100, 10, rl.White)
			rl.DrawText("borders on/off: / keypad", screenW-290, 120, 10, rl.White)
			rl.DrawText("debug on/off: del keypad", screenW-290, 140, 10, rl.White)
			rl.DrawText("left click to move", screenW-290, 160, 10, rl.White)

		}

		rl.EndDrawing()
	}

	rl.CloseWindow()

}
func getblock() { // MARK getblock
	addzoom := float32(1)
	if camera.Zoom == 2.0 {
		addzoom = 2
	} else if camera.Zoom == 4.0 {

	}
	xchange := float32(0)
	if mousepos.Y < 85*addzoom {
		for a := 0; a < 15; a++ {
			if mousepos.X < (150*addzoom)+xchange && mousepos.X > (50*addzoom)+xchange {
				blockactive = a
			}
			xchange += 100
		}
	}
	if mousepos.Y > 85*addzoom && mousepos.Y < 110*addzoom {
		for a := 0; a < 14; a++ {
			if mousepos.X < (150*addzoom)+xchange && mousepos.X > (50*addzoom)+xchange {
				blockactive = a + 15
			}
			xchange += 100
		}
	}
	if mousepos.Y > 110*addzoom && mousepos.Y < 135*addzoom {
		for a := 0; a < 15; a++ {
			if mousepos.X < (150*addzoom)+xchange && mousepos.X > (50*addzoom)+xchange {
				blockactive = a + 29
			}
			xchange += 100
		}
	}
	if mousepos.Y > 135*addzoom && mousepos.Y < 160*addzoom {
		for a := 0; a < 14; a++ {
			if mousepos.X < (150*addzoom)+xchange && mousepos.X > (50*addzoom)+xchange {
				blockactive = a + 44
			}
			xchange += 100
		}
	}
	if mousepos.Y > 160*addzoom && mousepos.Y < 185*addzoom {
		for a := 0; a < 15; a++ {
			if mousepos.X < (150*addzoom)+xchange && mousepos.X > (50*addzoom)+xchange {
				blockactive = a + 58
			}
			xchange += 100
		}
	}
	if mousepos.Y > 185*addzoom && mousepos.Y < 210*addzoom {
		for a := 0; a < 14; a++ {
			if mousepos.X < (150*addzoom)+xchange && mousepos.X > (50*addzoom)+xchange {
				blockactive = a + 73
			}
			xchange += 100
		}
	}
	if mousepos.Y > 210*addzoom && mousepos.Y < 235*addzoom {
		for a := 0; a < 15; a++ {
			if mousepos.X < (150*addzoom)+xchange && mousepos.X > (50*addzoom)+xchange {
				blockactive = a + 87
			}
			xchange += 100
		}
	}
	if mousepos.Y > 235*addzoom && mousepos.Y < 260*addzoom {
		for a := 0; a < 14; a++ {
			if mousepos.X < (150*addzoom)+xchange && mousepos.X > (50*addzoom)+xchange {
				blockactive = a + 102
			}
			xchange += 100
		}
	}
	if mousepos.Y > 260*addzoom && mousepos.Y < 285*addzoom {
		for a := 0; a < 15; a++ {
			if mousepos.X < (150*addzoom)+xchange && mousepos.X > (50*addzoom)+xchange {
				blockactive = a + 116
			}
			xchange += 100
		}
	}
	if mousepos.Y > 285*addzoom && mousepos.Y < 310*addzoom {
		for a := 0; a < 14; a++ {
			if mousepos.X < (150*addzoom)+xchange && mousepos.X > (50*addzoom)+xchange {
				blockactive = a + 131
			}
			xchange += 100
		}
	}
	if mousepos.Y > 310*addzoom && mousepos.Y < 335*addzoom {
		for a := 0; a < 15; a++ {
			if mousepos.X < (150*addzoom)+xchange && mousepos.X > (50*addzoom)+xchange {
				blockactive = a + 145
			}
			xchange += 100
		}
	}
	if mousepos.Y > 335*addzoom && mousepos.Y < 360*addzoom {
		for a := 0; a < 14; a++ {
			if mousepos.X < (150*addzoom)+xchange && mousepos.X > (50*addzoom)+xchange {
				blockactive = a + 160
			}
			xchange += 100
		}
	}
	if mousepos.Y > 360*addzoom && mousepos.Y < 385*addzoom {
		for a := 0; a < 15; a++ {
			if mousepos.X < (150*addzoom)+xchange && mousepos.X > (50*addzoom)+xchange {
				blockactive = a + 174
			}
			xchange += 100
		}
	}
	if mousepos.Y > 385*addzoom && mousepos.Y < 410*addzoom {
		for a := 0; a < 14; a++ {
			if mousepos.X < (150*addzoom)+xchange && mousepos.X > (50*addzoom)+xchange {
				blockactive = a + 189
			}
			xchange += 100
		}
	}
	if mousepos.Y > 410*addzoom && mousepos.Y < 435*addzoom {
		for a := 0; a < 15; a++ {
			if mousepos.X < (150*addzoom)+xchange && mousepos.X > (50*addzoom)+xchange {
				blockactive = a + 203
			}
			xchange += 100
		}
	}
	if mousepos.Y > 435*addzoom && mousepos.Y < 460*addzoom {
		for a := 0; a < 14; a++ {
			if mousepos.X < (150*addzoom)+xchange && mousepos.X > (50*addzoom)+xchange {
				blockactive = a + 218
			}
			xchange += 100
		}
	}
	if mousepos.Y > 460*addzoom && mousepos.Y < 485*addzoom {
		for a := 0; a < 15; a++ {
			if mousepos.X < (150*addzoom)+xchange && mousepos.X > (50*addzoom)+xchange {
				blockactive = a + 232
			}
			xchange += 100
		}
	}
	if mousepos.Y > 485*addzoom && mousepos.Y < 510*addzoom {
		for a := 0; a < 14; a++ {
			if mousepos.X < (150*addzoom)+xchange && mousepos.X > (50*addzoom)+xchange {
				blockactive = a + 247
			}
			xchange += 100
		}
	}
	if mousepos.Y > 510*addzoom && mousepos.Y < 535*addzoom {
		for a := 0; a < 15; a++ {
			if mousepos.X < (150*addzoom)+xchange && mousepos.X > (50*addzoom)+xchange {
				blockactive = a + 261
			}
			xchange += 100
		}
	}
	if mousepos.Y > 535*addzoom && mousepos.Y < 560*addzoom {
		for a := 0; a < 14; a++ {
			if mousepos.X < (150*addzoom)+xchange && mousepos.X > (50*addzoom)+xchange {
				blockactive = a + 276
			}
			xchange += 100
		}
	}
	if mousepos.Y > 560*addzoom && mousepos.Y < 585*addzoom {
		for a := 0; a < 15; a++ {
			if mousepos.X < (150*addzoom)+xchange && mousepos.X > (50*addzoom)+xchange {
				blockactive = a + 290
			}
			xchange += 100
		}
	}
	if mousepos.Y > 585*addzoom && mousepos.Y < 610*addzoom {
		for a := 0; a < 14; a++ {
			if mousepos.X < (150*addzoom)+xchange && mousepos.X > (50*addzoom)+xchange {
				blockactive = a + 305
			}
			xchange += 100
		}
	}
	if mousepos.Y > 610*addzoom && mousepos.Y < 635*addzoom {
		for a := 0; a < 15; a++ {
			if mousepos.X < (150*addzoom)+xchange && mousepos.X > (50*addzoom)+xchange {
				blockactive = a + 319
			}
			xchange += 100
		}
	}
	if mousepos.Y > 635*addzoom && mousepos.Y < 660*addzoom {
		for a := 0; a < 14; a++ {
			if mousepos.X < (150*addzoom)+xchange && mousepos.X > (50*addzoom)+xchange {
				blockactive = a + 334
			}
			xchange += 100
		}
	}
	if mousepos.Y > 660*addzoom && mousepos.Y < 685*addzoom {
		for a := 0; a < 15; a++ {
			if mousepos.X < (150*addzoom)+xchange && mousepos.X > (50*addzoom)+xchange {
				blockactive = a + 348
			}
			xchange += 100
		}
	}
	if mousepos.Y > 685*addzoom && mousepos.Y < 710*addzoom {
		for a := 0; a < 14; a++ {
			if mousepos.X < (150*addzoom)+xchange && mousepos.X > (50*addzoom)+xchange {
				blockactive = a + 363
			}
			xchange += 100
		}
	}
	if mousepos.Y > 710*addzoom && mousepos.Y < 735*addzoom {
		for a := 0; a < 15; a++ {
			if mousepos.X < (150*addzoom)+xchange && mousepos.X > (50*addzoom)+xchange {
				blockactive = a + 377
			}
			xchange += 100
		}
	}
	if mousepos.Y > 735*addzoom && mousepos.Y < 760*addzoom {
		for a := 0; a < 14; a++ {
			if mousepos.X < (150*addzoom)+xchange && mousepos.X > (50*addzoom)+xchange {
				blockactive = a + 392
			}
			xchange += 100
		}
	}
	if mousepos.Y > 760*addzoom && mousepos.Y < 785*addzoom {
		for a := 0; a < 15; a++ {
			if mousepos.X < (150*addzoom)+xchange && mousepos.X > (50*addzoom)+xchange {
				blockactive = a + 406
			}
			xchange += 100
		}
	}
	if mousepos.Y > 785*addzoom && mousepos.Y < 810*addzoom {
		for a := 0; a < 14; a++ {
			if mousepos.X < (150*addzoom)+xchange && mousepos.X > (50*addzoom)+xchange {
				blockactive = a + 421
			}
			xchange += 100
		}
	}
	if mousepos.Y > 810*addzoom && mousepos.Y < 835*addzoom {
		for a := 0; a < 15; a++ {
			if mousepos.X < (150*addzoom)+xchange && mousepos.X > (50*addzoom)+xchange {
				blockactive = a + 435
			}
			xchange += 100
		}
	}
}

// random numbers
func rInt(min, max int) int {
	return rand.Intn(max-min) + min
}
func rInt32(min, max int) int32 {
	a := int32(rand.Intn(max-min) + min)
	return a
}
func rFloat32(min, max int) float32 {
	a := float32(rand.Intn(max-min) + min)
	return a
}
func flipcoin() bool {
	var b bool
	a := rInt(0, 10001)
	if a < 5000 {
		b = true
	}
	return b
}
func rolldice() int {
	a := rInt(1, 7)
	return a
}
