run:
	(cd backend && go run main.go) & (cd frontend && npm run dev)

init:
	cd frontend && npm install
	cd backend && go mod tidy
