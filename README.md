# MUMAX3-CQED: mumax<sup>3</sup> Cavity QED

This software adds a new feature to open source software called [mumax<sup>3</sup>](https://mumax.github.io/). The new MUMAX3-CQED is also open source and it is available at [https://github.com/sergiomtzlosa/mumax3-cqed](https://github.com/sergiomtzlosa/mumax3-cqed).

The original mumax<sup>3</sup> code contains lots of files managing the frontend and the CUDA calculations, and for this purpose we list below the files we modified or created to develop our new feature:

```bash
$ 📦src/github.com/mumax/3
  ├── 📂cmd
  │   └─ 📂mumax3
  │   	 ├── 📄mumax3.exe (modified file)
  │   	 └── 📄main.go (modified file)
  ├── 📂cuda
  │   ├── 📄calcspincbeff.cu (new file)
  │   ├── 📄calcspinbeffdissipation.cu (new file)
  │   ├── 📄lltorque.go (modified file)
  │   ├── 📄Makefile (modified file)
  │   ├── 📄make.ps1 (new file)
  │   └── 📄realclean.ps1 (new file)
  └── 📂engine
      ├── 📄run.go (modified file)
      ├── 📄torque.go (modified file)
      ├── 📄effectivefield.go (modified file)
      ├── 📄utils_extension.go (new file)
      └── 📄bib.go (modified file)
```

The files under `cuda` folder manage the operations in the GPU and files below `engine` folder manage the input/output data from/to the GPU and also present the data to the user.

To run the mumax3-cqed binary, open a shell and run your script as:
```console
$ mumax.exe script-file.mx3
```
or in *NIX systems:
```console
$ mumax script-file.mx3
```

#### INSTALLATION IN *NIX SYSTEMS
---------------------------------

It is required to install [CUDA Toolkit](https://developer.nvidia.com/cuda-downloads). This mumax3-cqed software has been tested in Windows 10 with [CUDA Toolkit 10.2](https://developer.nvidia.com/cuda-10.2-download-archive) and Debian GNU/Linux (Debian 12 Bookworm) with [CUDA Toolkit 12.0](https://developer.nvidia.com/cuda-12-0-0-download-archive), it is also required to install git. Follow [this guide](https://www.server-world.info/en/note?os=Debian_12&p=nvidia&f=1) to install CUDA 12.0 in Debian 12 Bookworm.

The installation process on *NIX systems takes place with the following commands:

```bash
$ cd $HOME
$ sudo apt update && sudo apt install git
$ curl -OL https://golang.google.cn/dl/go1.9.linux-amd64.tar.gz
$ tar -xf go1.9.linux-amd64.tar.gz
$ mv go go1.9
$ export PATH=$(pwd)/go1.9/bin:$PATH
$ export GOPATH=$(pwd)/go
$ mkdir -p go/src/github.com/mumax/3
$ git clone https://github.com/sergiomtzlosa/mumax3-cqed go/src/github.com/mumax/3
$ cd go/src/github.com/mumax/3/cuda    
$ make
$ cd go/src/github.com/mumax/3/cmd/mumax3
$ go install -v "github.com/mumax/3/..."
$ export PATH=$HOME/go/bin:$PATH
```

The binary file is placed in `$HOME/go/bin`

#### INSTALLATION IN WINDOWS SYSTEMS
------------------------------------

For Windows systems install [Visual Studio Community 2019](https://visualstudio.microsoft.com/en/vs/older-downloads/) with **.Net desktop development**, **Desktop development with C++**, **Universal Windows Platform development**.

![vs2019-community-installer](./images/vs2019-community-installer.png)

Once the installation of Visual Studio Community 2019 finishes, put the following in the Path variable:

```text
c:\Program Files (x86)\Microsoft Visual Studio\2019\Community\VC\Tools\MSVC\14.29.30133\bin\Hostx64\x64
```

Install [CUDA Tollkit 10.2](https://developer.nvidia.com/cuda-10.2-download-archive) (or later). Follow these steps to install mumax3-cqed in Windows 10:

1. Install [Git for Windows](https://git-scm.com/download/win)
2. Download [Golang 1.9](https://dl.google.com/go/go1.9.windows-amd64.zip) and uncompress the file in `c:\`, rename the folder to go1.9: `c:\go1.9`
3. Add to Path enviroment variable the golang binaries path: `c:\go1.9\bin`
4. Set a environment variable called `GOPATH` poiting to `c:\go` and create the folder in your system, this folder will be your working path for the source code
5. Run in PowerShell window: `mkdir -p c:\go\src\github.com\mumax\3`
6. Clone the repository: `git clone https://github.com/sergiomtzlosa/mumax3-cqed c:\go\src\github.com\mumax\3`
7. Open a PowerShell window and type: `cd c:\go\src\github.com\mumax\3\cuda`
8. Compile the CUDA files with the following command: `.\make.ps1`
9. Navigate to binary folder: `cd c:\go\src\github.com\mumax\3\cmd\mumax3`
10. Compile mumax3-cqed main binary file: `go install -v "github.com/mumax/3/..."`

Alternatively, you can use Windows Subsystem Linux (WSL) just to compile CUDA code:

1. Install [CUDA Toolkit 10.2](https://developer.nvidia.com/cuda-10.2-download-archive) for WSL
2. Open a Windows Subsystem Linux terminal and navigate to cuda files folder: `cd /mnt/c/go/src/github.com/mumax/3/cuda`
3. Compile cuda files in WSL: `make`

The binary file is placed in `c:\go\bin`

Add `c:\go\bin` to Path to call the `mumax.exe` binary file from shell
