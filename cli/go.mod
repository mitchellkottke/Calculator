module calculator/calculator

go 1.22.5

replace (
	calculator/calculations => ../calculations
	calculator/queue => ../queue
	calculator/stack => ../stack
)

require calculator/calculations v0.0.0-00010101000000-000000000000

require (
	calculator/stack v0.0.0-00010101000000-000000000000 // indirect
)
