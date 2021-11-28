package main

import spx "github.com/goplus/spx"

var niuPlay bool = false
var niuFlap bool = true
var niuDie bool = false
var niuGravity float64 = 0.0
var gameScrollX float64 = -3.5
var gameScrollTime float64 = 0.05

type index struct {
	spx.Game
	Logo      Logo
	GameOver  GameOver
	Ground    Ground
	Niu       Niu
	Pipe      Pipe
	Start     Start
	TapInfo   TapInfo
	userScore int
}
type Logo struct {
	spx.Sprite
	*index
}
type GameOver struct {
	spx.Sprite
	*index
}
type Ground struct {
	spx.Sprite
	*index
}
type Niu struct {
	spx.Sprite
	*index
}
type Pipe struct {
	spx.Sprite
	*index
}
type Start struct {
	spx.Sprite
	*index
}
type TapInfo struct {
	spx.Sprite
	*index
}

func (this *index) MainEntry() {
//line /Users/detu/github/goplus/FlappyBird/index.gmx:21
	spx.Gopt_Game_Run(this, "res", &spx.Config{Title: "FlappyBird (by Go+)"})
}
func main() {
	spx.Gopt_Game_Main(new(index))
}
func (this *Ground) Main() {
//line /Users/detu/github/goplus/FlappyBird/Ground.spx:1
	this.OnStart(func() {
//line /Users/detu/github/goplus/FlappyBird/Ground.spx:2
		this.SetXYpos(0, -130)
//line /Users/detu/github/goplus/FlappyBird/Ground.spx:3
		this.Show()
	})
}
func (this *Logo) Main() {
//line /Users/detu/github/goplus/FlappyBird/Logo.spx:1
	this.OnStart(func() {
//line /Users/detu/github/goplus/FlappyBird/Logo.spx:2
		for {
			spx.Sched()
//line /Users/detu/github/goplus/FlappyBird/Logo.spx:3
			if niuPlay == true {
//line /Users/detu/github/goplus/FlappyBird/Logo.spx:4
				return
			}
//line /Users/detu/github/goplus/FlappyBird/Logo.spx:6
			if niuDie == false {
//line /Users/detu/github/goplus/FlappyBird/Logo.spx:7
				for
//line /Users/detu/github/goplus/FlappyBird/Logo.spx:7
				i := 0; i < 10;
//line /Users/detu/github/goplus/FlappyBird/Logo.spx:7
				i++ {
					spx.Sched()
//line /Users/detu/github/goplus/FlappyBird/Logo.spx:8
					if niuPlay == false {
//line /Users/detu/github/goplus/FlappyBird/Logo.spx:9
						this.ChangeYpos(-0.5)
					}
//line /Users/detu/github/goplus/FlappyBird/Logo.spx:11
					this.Wait(0.05)
				}
//line /Users/detu/github/goplus/FlappyBird/Logo.spx:14
				for
//line /Users/detu/github/goplus/FlappyBird/Logo.spx:14
				i := 0; i < 10;
//line /Users/detu/github/goplus/FlappyBird/Logo.spx:14
				i++ {
					spx.Sched()
//line /Users/detu/github/goplus/FlappyBird/Logo.spx:15
					if niuPlay == false {
//line /Users/detu/github/goplus/FlappyBird/Logo.spx:16
						this.ChangeYpos(0.5)
					}
//line /Users/detu/github/goplus/FlappyBird/Logo.spx:18
					this.Wait(0.05)
				}
			}
		}
	})
//line /Users/detu/github/goplus/FlappyBird/Logo.spx:25
	this.OnMsg__1("menu", func() {
//line /Users/detu/github/goplus/FlappyBird/Logo.spx:26
		this.SetXYpos(-23, 55)
//line /Users/detu/github/goplus/FlappyBird/Logo.spx:27
		this.GotoFront()
//line /Users/detu/github/goplus/FlappyBird/Logo.spx:28
		this.Show()
	})
//line /Users/detu/github/goplus/FlappyBird/Logo.spx:31
	this.OnMsg__1("start", func() {
//line /Users/detu/github/goplus/FlappyBird/Logo.spx:32
		this.Hide()
	})
//line /Users/detu/github/goplus/FlappyBird/Logo.spx:35
	this.OnMsg__1("tap", func() {
//line /Users/detu/github/goplus/FlappyBird/Logo.spx:36
		this.Hide()
	})
}
func (this *Niu) Main() {
//line /Users/detu/github/goplus/FlappyBird/Niu.spx:1
	this.OnMsg__1("start", func() {
//line /Users/detu/github/goplus/FlappyBird/Niu.spx:2
		this.SetXYpos(-80, 0)
	})
//line /Users/detu/github/goplus/FlappyBird/Niu.spx:5
	this.OnMsg__1("menu", func() {
//line /Users/detu/github/goplus/FlappyBird/Niu.spx:6
		this.SetXYpos(100.0, 55.0)
//line /Users/detu/github/goplus/FlappyBird/Niu.spx:7
		this.TurnTo(90)
//line /Users/detu/github/goplus/FlappyBird/Niu.spx:8
		this.Wait(0.25)
//line /Users/detu/github/goplus/FlappyBird/Niu.spx:9
		for {
			spx.Sched()
//line /Users/detu/github/goplus/FlappyBird/Niu.spx:10
			if niuDie == true {
//line /Users/detu/github/goplus/FlappyBird/Niu.spx:11
				this.NextCostume()
//line /Users/detu/github/goplus/FlappyBird/Niu.spx:12
				this.Wait(0.1)
			}
//line /Users/detu/github/goplus/FlappyBird/Niu.spx:14
			this.Wait(0.03)
		}
	})
//line /Users/detu/github/goplus/FlappyBird/Niu.spx:19
	this.OnMsg__1("menu", func() {
//line /Users/detu/github/goplus/FlappyBird/Niu.spx:21
		for {
			spx.Sched()
//line /Users/detu/github/goplus/FlappyBird/Niu.spx:22
			if niuPlay == true {
//line /Users/detu/github/goplus/FlappyBird/Niu.spx:23
				return
			}
//line /Users/detu/github/goplus/FlappyBird/Niu.spx:25
			if niuDie == false {
//line /Users/detu/github/goplus/FlappyBird/Niu.spx:26
				for
//line /Users/detu/github/goplus/FlappyBird/Niu.spx:26
				i := 0; i < 10;
//line /Users/detu/github/goplus/FlappyBird/Niu.spx:26
				i++ {
					spx.Sched()
//line /Users/detu/github/goplus/FlappyBird/Niu.spx:27
					if niuPlay == false {
//line /Users/detu/github/goplus/FlappyBird/Niu.spx:28
						this.ChangeYpos(-0.5)
					}
//line /Users/detu/github/goplus/FlappyBird/Niu.spx:30
					this.Wait(0.05)
				}
//line /Users/detu/github/goplus/FlappyBird/Niu.spx:33
				for
//line /Users/detu/github/goplus/FlappyBird/Niu.spx:33
				i := 0; i < 10;
//line /Users/detu/github/goplus/FlappyBird/Niu.spx:33
				i++ {
					spx.Sched()
//line /Users/detu/github/goplus/FlappyBird/Niu.spx:34
					if niuPlay == false {
//line /Users/detu/github/goplus/FlappyBird/Niu.spx:35
						this.ChangeYpos(0.5)
					}
//line /Users/detu/github/goplus/FlappyBird/Niu.spx:37
					this.Wait(0.05)
				}
			}
		}
	})
//line /Users/detu/github/goplus/FlappyBird/Niu.spx:46
	this.OnMsg__1("tap", func() {
//line /Users/detu/github/goplus/FlappyBird/Niu.spx:48
		for {
			spx.Sched()
//line /Users/detu/github/goplus/FlappyBird/Niu.spx:49
			if niuPlay == true {
//line /Users/detu/github/goplus/FlappyBird/Niu.spx:50
				niuGravity = niuGravity - 0.25
//line /Users/detu/github/goplus/FlappyBird/Niu.spx:51
				this.Wait(0.05)
			}
//line /Users/detu/github/goplus/FlappyBird/Niu.spx:53
			this.Wait(0.01)
		}
	})
//line /Users/detu/github/goplus/FlappyBird/Niu.spx:57
	this.OnMsg__1("tap", func() {
//line /Users/detu/github/goplus/FlappyBird/Niu.spx:59
		for {
			spx.Sched()
//line /Users/detu/github/goplus/FlappyBird/Niu.spx:60
			if niuPlay == true {
//line /Users/detu/github/goplus/FlappyBird/Niu.spx:62
				for !(this.KeyPressed(spx.KeySpace) || this.MousePressed()) {
					spx.Sched()
//line /Users/detu/github/goplus/FlappyBird/Niu.spx:63
					this.Wait(0.01)
				}
//line /Users/detu/github/goplus/FlappyBird/Niu.spx:65
				if niuDie == false {
//line /Users/detu/github/goplus/FlappyBird/Niu.spx:66
					niuGravity = 0.8
//line /Users/detu/github/goplus/FlappyBird/Niu.spx:67
					for
//line /Users/detu/github/goplus/FlappyBird/Niu.spx:67
					i := 0; i < 10;
//line /Users/detu/github/goplus/FlappyBird/Niu.spx:67
					i++ {
						spx.Sched()
//line /Users/detu/github/goplus/FlappyBird/Niu.spx:68
						this.ChangeYpos(3.5)
//line /Users/detu/github/goplus/FlappyBird/Niu.spx:69
						this.Wait(0.03)
					}
				}
			}
//line /Users/detu/github/goplus/FlappyBird/Niu.spx:74
			this.Wait(0.01)
		}
	})
//line /Users/detu/github/goplus/FlappyBird/Niu.spx:78
	this.OnMsg__1("tap", func() {
//line /Users/detu/github/goplus/FlappyBird/Niu.spx:80
		for {
			spx.Sched()
//line /Users/detu/github/goplus/FlappyBird/Niu.spx:82
			if this.Touching("Pipe") || this.Ypos() < -123.9 {
//line /Users/detu/github/goplus/FlappyBird/Niu.spx:83
				niuPlay = false
//line /Users/detu/github/goplus/FlappyBird/Niu.spx:84
				niuDie = true
//line /Users/detu/github/goplus/FlappyBird/Niu.spx:85
				niuGravity = 0
//line /Users/detu/github/goplus/FlappyBird/Niu.spx:87
				this.Broadcast__0("flash")
//line /Users/detu/github/goplus/FlappyBird/Niu.spx:88
				this.Broadcast__0("game over")
//line /Users/detu/github/goplus/FlappyBird/Niu.spx:89
				this.GotoFront()
//line /Users/detu/github/goplus/FlappyBird/Niu.spx:91
				break
			}
//line /Users/detu/github/goplus/FlappyBird/Niu.spx:94
			this.Wait(0.03)
		}
	})
//line /Users/detu/github/goplus/FlappyBird/Niu.spx:98
	var flag int = 0
//line /Users/detu/github/goplus/FlappyBird/Niu.spx:101
	this.OnMsg__1("tap", func() {
//line /Users/detu/github/goplus/FlappyBird/Niu.spx:102
		for {
			spx.Sched()
//line /Users/detu/github/goplus/FlappyBird/Niu.spx:103
			if niuPlay == true && niuGravity != 0.0 {
//line /Users/detu/github/goplus/FlappyBird/Niu.spx:104
				if this.Heading() < 150 && flag == 0 {
//line /Users/detu/github/goplus/FlappyBird/Niu.spx:105
					this.Turn(niuGravity * -0.75)
				} else {
//line /Users/detu/github/goplus/FlappyBird/Niu.spx:107
					if this.Heading() < 90 {
//line /Users/detu/github/goplus/FlappyBird/Niu.spx:108
						flag = 0
					} else {
//line /Users/detu/github/goplus/FlappyBird/Niu.spx:110
						flag = 1
					}
//line /Users/detu/github/goplus/FlappyBird/Niu.spx:112
					this.Turn(-(niuGravity * -0.75))
				}
//line /Users/detu/github/goplus/FlappyBird/Niu.spx:114
				this.ChangeYpos(niuGravity)
			}
//line /Users/detu/github/goplus/FlappyBird/Niu.spx:116
			this.Wait(0.03)
		}
	})
//line /Users/detu/github/goplus/FlappyBird/Niu.spx:122
	this.OnMsg__1("game over", func() {
//line /Users/detu/github/goplus/FlappyBird/Niu.spx:124
		for {
			spx.Sched()
//line /Users/detu/github/goplus/FlappyBird/Niu.spx:125
			if !this.Touching("Ground") {
//line /Users/detu/github/goplus/FlappyBird/Niu.spx:126
				this.ChangeYpos(-10)
			}
//line /Users/detu/github/goplus/FlappyBird/Niu.spx:128
			this.Wait(0.04)
		}
//line /Users/detu/github/goplus/FlappyBird/Niu.spx:130
		this.SetYpos(-124)
//line /Users/detu/github/goplus/FlappyBird/Niu.spx:131
		this.TurnTo(180)
	})
}
func (this *Pipe) Main() {
//line /Users/detu/github/goplus/FlappyBird/Pipe.spx:1
	this.OnStart(func() {
//line /Users/detu/github/goplus/FlappyBird/Pipe.spx:2
		this.Hide()
//line /Users/detu/github/goplus/FlappyBird/Pipe.spx:3
		this.SetXYpos(0, 0)
	})
//line /Users/detu/github/goplus/FlappyBird/Pipe.spx:6
	this.OnMsg__1("tap", func() {
//line /Users/detu/github/goplus/FlappyBird/Pipe.spx:7
		this.Wait(1.5)
//line /Users/detu/github/goplus/FlappyBird/Pipe.spx:8
		for {
			spx.Sched()
//line /Users/detu/github/goplus/FlappyBird/Pipe.spx:9
			if niuDie == false {
//line /Users/detu/github/goplus/FlappyBird/Pipe.spx:10
				spx.Gopt_Sprite_Clone__0(this)
//line /Users/detu/github/goplus/FlappyBird/Pipe.spx:11
				this.Wait(2.5)
			}
//line /Users/detu/github/goplus/FlappyBird/Pipe.spx:13
			this.Wait(0.03)
		}
	})
//line /Users/detu/github/goplus/FlappyBird/Pipe.spx:17
	this.OnCloned__1(func() {
//line /Users/detu/github/goplus/FlappyBird/Pipe.spx:18
		this.SetCostume(spx.Rand__0(0, 4))
//line /Users/detu/github/goplus/FlappyBird/Pipe.spx:19
		this.SetXYpos(250, 17)
//line /Users/detu/github/goplus/FlappyBird/Pipe.spx:20
		this.Show()
//line /Users/detu/github/goplus/FlappyBird/Pipe.spx:21
		for {
			spx.Sched()
//line /Users/detu/github/goplus/FlappyBird/Pipe.spx:22
			if this.Xpos() < -250 {
//line /Users/detu/github/goplus/FlappyBird/Pipe.spx:23
				break
			}
//line /Users/detu/github/goplus/FlappyBird/Pipe.spx:25
			if niuDie == false {
//line /Users/detu/github/goplus/FlappyBird/Pipe.spx:26
				this.ChangeXpos(gameScrollX)
			}
//line /Users/detu/github/goplus/FlappyBird/Pipe.spx:29
			this.Wait(gameScrollTime)
		}
//line /Users/detu/github/goplus/FlappyBird/Pipe.spx:31
		this.Die()
	})
//line /Users/detu/github/goplus/FlappyBird/Pipe.spx:35
	this.OnCloned__1(func() {
//line /Users/detu/github/goplus/FlappyBird/Pipe.spx:36
		for {
			spx.Sched()
//line /Users/detu/github/goplus/FlappyBird/Pipe.spx:37
			if this.Xpos() < -80.0 && niuPlay == true {
//line /Users/detu/github/goplus/FlappyBird/Pipe.spx:38
				this.userScore++
//line /Users/detu/github/goplus/FlappyBird/Pipe.spx:39
				return
			}
//line /Users/detu/github/goplus/FlappyBird/Pipe.spx:41
			this.Wait(0.03)
		}
	})
//line /Users/detu/github/goplus/FlappyBird/Pipe.spx:45
	this.OnMsg__1("menu", func() {
//line /Users/detu/github/goplus/FlappyBird/Pipe.spx:46
		this.Die()
	})
}
func (this *Start) Main() {
//line /Users/detu/github/goplus/FlappyBird/Start.spx:1
	this.OnStart(func() {
//line /Users/detu/github/goplus/FlappyBird/Start.spx:2
		this.Broadcast__0("menu")
//line /Users/detu/github/goplus/FlappyBird/Start.spx:3
		this.SetXYpos(0, -95)
	})
//line /Users/detu/github/goplus/FlappyBird/Start.spx:6
	this.OnStart(func() {
//line /Users/detu/github/goplus/FlappyBird/Start.spx:7
		for {
			spx.Sched()
//line /Users/detu/github/goplus/FlappyBird/Start.spx:8
			if this.Visible() == false {
//line /Users/detu/github/goplus/FlappyBird/Start.spx:9
				break
			}
//line /Users/detu/github/goplus/FlappyBird/Start.spx:11
			if this.Touching(spx.Mouse) {
//line /Users/detu/github/goplus/FlappyBird/Start.spx:12
				this.SetYpos(-90)
			} else {
//line /Users/detu/github/goplus/FlappyBird/Start.spx:14
				this.SetYpos(-95)
			}
//line /Users/detu/github/goplus/FlappyBird/Start.spx:16
			this.Wait(0.1)
		}
	})
//line /Users/detu/github/goplus/FlappyBird/Start.spx:20
	this.OnClick(func() {
//line /Users/detu/github/goplus/FlappyBird/Start.spx:21
		this.Broadcast__0("start")
//line /Users/detu/github/goplus/FlappyBird/Start.spx:22
		this.Hide()
	})
//line /Users/detu/github/goplus/FlappyBird/Start.spx:25
	this.OnMsg__1("start", func() {
//line /Users/detu/github/goplus/FlappyBird/Start.spx:26
		this.Hide()
	})
//line /Users/detu/github/goplus/FlappyBird/Start.spx:28
	this.OnMsg__1("menu", func() {
//line /Users/detu/github/goplus/FlappyBird/Start.spx:29
		this.GotoFront()
//line /Users/detu/github/goplus/FlappyBird/Start.spx:30
		this.Show()
	})
}
func (this *TapInfo) Main() {
//line /Users/detu/github/goplus/FlappyBird/TapInfo.spx:3
	this.OnStart(func() {
//line /Users/detu/github/goplus/FlappyBird/TapInfo.spx:4
		this.SetXYpos(0, 0)
//line /Users/detu/github/goplus/FlappyBird/TapInfo.spx:5
		this.Hide()
	})
//line /Users/detu/github/goplus/FlappyBird/TapInfo.spx:8
	this.OnMsg__1("start", func() {
//line /Users/detu/github/goplus/FlappyBird/TapInfo.spx:9
		this.Show()
//line /Users/detu/github/goplus/FlappyBird/TapInfo.spx:10
		this.Wait(0.25)
//line /Users/detu/github/goplus/FlappyBird/TapInfo.spx:11
		for !(this.KeyPressed(spx.KeySpace) || this.MousePressed()) {
			spx.Sched()
//line /Users/detu/github/goplus/FlappyBird/TapInfo.spx:12
			this.Wait(0.03)
		}
//line /Users/detu/github/goplus/FlappyBird/TapInfo.spx:14
		niuPlay = true
//line /Users/detu/github/goplus/FlappyBird/TapInfo.spx:15
		this.Broadcast__0("tap")
//line /Users/detu/github/goplus/FlappyBird/TapInfo.spx:16
		this.Hide()
	})
//line /Users/detu/github/goplus/FlappyBird/TapInfo.spx:20
	this.OnMsg__1("menu", func() {
//line /Users/detu/github/goplus/FlappyBird/TapInfo.spx:21
		this.Hide()
	})
}
func (this *GameOver) Main() {
//line /Users/detu/github/goplus/FlappyBird/GameOver.spx:1
	this.OnStart(func() {
//line /Users/detu/github/goplus/FlappyBird/GameOver.spx:2
		this.SetXYpos(0, -220)
//line /Users/detu/github/goplus/FlappyBird/GameOver.spx:3
		this.Hide()
	})
//line /Users/detu/github/goplus/FlappyBird/GameOver.spx:5
	this.OnMsg__1("game over", func() {
//line /Users/detu/github/goplus/FlappyBird/GameOver.spx:6
		this.Wait(0.5)
//line /Users/detu/github/goplus/FlappyBird/GameOver.spx:7
		this.GotoFront()
//line /Users/detu/github/goplus/FlappyBird/GameOver.spx:8
		this.Show()
//line /Users/detu/github/goplus/FlappyBird/GameOver.spx:9
		this.Glide__0(0, 0, 0.25)
	})
//line /Users/detu/github/goplus/FlappyBird/GameOver.spx:12
	this.OnMsg__1("menu", func() {
//line /Users/detu/github/goplus/FlappyBird/GameOver.spx:13
		this.Hide()
//line /Users/detu/github/goplus/FlappyBird/GameOver.spx:14
		this.SetXYpos(0, -220)
	})
}
