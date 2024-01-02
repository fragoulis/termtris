package main

type Shape struct {
	Frames []Frame
}

var shapes = []Shape{
	Shape{
		Frames: []Frame{
			Frame{
				Width:  3,
				Height: 2,
				Data: []rune{
					0, '█', '█',
					'█', '█', 0,
				},
			},
			Frame{
				Width:  2,
				Height: 3,
				Data: []rune{
					'█', 0,
					'█', '█',
					0, '█',
				},
			},
			Frame{
				Width:  3,
				Height: 2,
				Data: []rune{
					'█', '█', 0,
					0, '█', '█',
				},
			},
			Frame{
				Width:  2,
				Height: 3,
				Data: []rune{
					0, '█',
					'█', '█',
					'█', 0,
				},
			},
		},
	},
	Shape{
		Frames: []Frame{
			Frame{
				Width:  3,
				Height: 2,
				Data: []rune{
					0, '█', 0,
					'█', '█', '█',
				},
			},
			Frame{
				Width:  2,
				Height: 3,
				Data: []rune{
					'█', 0,
					'█', '█',
					'█', 0,
				},
			},
			Frame{
				Width:  3,
				Height: 2,
				Data: []rune{
					'█', '█', '█',
					0, '█', 0,
				},
			},
			Frame{
				Width:  2,
				Height: 3,
				Data: []rune{
					0, '█',
					'█', '█',
					0, '█',
				},
			},
		},
	},
	Shape{
		Frames: []Frame{
			Frame{
				Width:  3,
				Height: 2,
				Data: []rune{
					'█', 0, 0,
					'█', '█', '█',
				},
			},
			Frame{
				Width:  2,
				Height: 3,
				Data: []rune{
					'█', '█',
					'█', 0,
					'█', 0,
				},
			},
			Frame{
				Width:  3,
				Height: 2,
				Data: []rune{
					'█', '█', '█',
					0, 0, '█',
				},
			},
			Frame{
				Width:  2,
				Height: 3,
				Data: []rune{
					0, '█',
					0, '█',
					'█', '█',
				},
			},
		},
	},
	Shape{
		Frames: []Frame{
			Frame{
				Width:  4,
				Height: 1,
				Data: []rune{
					'█', '█', '█', '█',
				},
			},
			Frame{
				Width:  1,
				Height: 4,
				Data: []rune{
					'█',
					'█',
					'█',
					'█',
				},
			},
		},
	},
	Shape{
		Frames: []Frame{
			Frame{
				Width:  2,
				Height: 2,
				Data: []rune{
					'█', '█',
					'█', '█',
				},
			},
		},
	},
}
