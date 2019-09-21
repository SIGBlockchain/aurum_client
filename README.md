<img src = "assets/aurum_logo_readme.jpg"  alt="drawing" width="200"/>

## Installation
Please select the branch that reflects your OS.
[Linux Binary](https://github.com/SIGBlockchain/aurum_client/tree/linux_release)
[Mac OS X Binary](https://github.com/SIGBlockchain/aurum_client/tree/mac_release)
[Windows Binary](https://github.com/SIGBlockchain/aurum_client/tree/windows_release)
Navigate to the bright green button on the right hand side that says clone or download. Click download zip.
Find the zip file you downloaded and unzip to your desired directoy.
Open a console or terminal window and navigate to the the aurum_client directory
Type the following (if on a Linux / Mac OS X system - for Windows, please change the listed files to be executable and run):
sudo chmod +x setup.sh
sudo ./setup.sh
You should see the following message:
**ascii logo**
Welcome! Let's get everything set up.
can't load package: package .: no Go files in /Users/zlabaschin/.go/src/github.com/SIGBlockchain/aurum_client
Initializing Aurum wallet...
Wallet setup complete.
Okay your wallet's ready.
Run ./client -h or ./client --help to see all commands
Check out the wiki on the github page for more info: https://github.com/SIGBlockchain/aurum_client/wiki

A file called aurum_wallet.json was created. Copy this file and store it in another directory as a backup (not in the aurum directory). __It contains your private key. If you lose you private key, you will lose access to your wallet and it is not recoverable.__
You are now ready to use aurum. If you type ./client -h, you will see the following:
Usage of ./client:
  -info
    	wallet info
  -send string
    	value to send
  -setup
    	set up client
  -to string
    	recipient
  -update
    	update wallet info
  -version
    	client version

string means put the address of the person you want to interact with.
setup is to create a __new wallet__
update gets the latest transaction data

Aurum Client
=============
Aurum is a proprietary UIC ACM blockchain project. The current planned use case is a token to be exchanged among students. The aurum client is for accessing the Aurum network. If you wish to be a producer, please visit [Project Aurum](https://github.com/SIGBlockchain/project_aurum) for the producer application.

## Contribute
If you would like to contribute, please comment on an issue you'd like to take on. Then, make a branch based on `dev`. Once you've completed the issue make a pull request from your branch to `dev`. If you have any questions simply ask in a comment on the issue.

## Credits
- First and foremost a big thank you to everyone who has worked on the SIG Blockchain team, both past and present. 
- Thank you to the [Association of Computing Machinery UIC Chapter](https://acm.cs.uic.edu/), for providing the support that SIG Blockchain has needed to thrive.
- Shoutout to [calvinmorett](https://github.com/calvinmorett) for providing us with the awesome Aurum logo.

## License
MIT License
- See the [LICENSE](https://github.com/SIGBlockchain/project_aurum/blob/readme/LICENSE) file for details.
