# resume-go

This is a crude / basic implementation of jsonresume in Go. It is not compatible with jsonresume themes, the themes need to be adapted to Go `html/template`'s

## Usage via docker

- Create directory to work in  
`mkdir resume && cd resume`
- Initialise a new resume.json  
`docker run -v ${PWD}:/app techarchitect/resume-go:latest init`
- Pull a theme/template  
`git clone https://github.com/Tech-Arch1tect/go-jsonresume-theme-stackoverflow.git`
- Make your changes to resume.json  
- Generate HTML resume  
`docker run -v ${PWD}:/app techarchitect/resume-go:latest generate -t go-jsonresume-theme-stackoverflow`
- Open resume.html

## Usage via binary

- Download the binary from a release into your PATH
- Create directory to work in  
`mkdir resume && cd resume`
- Initialise a new resume.json  
`resume-go init`
- Pull a theme/template  
`git clone https://github.com/Tech-Arch1tect/go-jsonresume-theme-stackoverflow.git`
- Make your changes to resume.json  
- Generate HTML resume  
`resume-go generate -t go-jsonresume-theme-stackoverflow`
- Open resume.html