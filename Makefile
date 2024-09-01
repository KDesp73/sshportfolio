all: portfolio sshportfolio

portfolio:
	go build ./cmd/portfolio/portfolio.go

sshportfolio:
	go build ./cmd/sshportfolio/sshportfolio.go

clean:
	rm portfolio
	rm sshportfolio
