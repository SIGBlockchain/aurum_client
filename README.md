<img src = "assets/aurum_logo_readme.jpg"  alt="drawing" width="200"/>

## Installation
Please select the branch that reflects your OS.<br />
* [Linux Binary](https://github.com/SIGBlockchain/aurum_client/tree/linux_release.zip)
* [Mac OS X Binary](https://github.com/SIGBlockchain/aurum_client/tree/mac_release.zip)
* [Windows Binary](https://github.com/SIGBlockchain/aurum_client/tree/windows_release.zip)

<br />Navigate to the bright green button on the right hand side that says clone or download. Click download zip.<br />
Find the zip file you downloaded and unzip to your desired directoy.<br />
Open a console or terminal window and navigate to the the aurum_client directory<br />
Type the following (for a Linux / Mac OS X environment): <br />
1. sudo chmod +x setup.sh
2. sudo ./setup.sh

* For Windows, please change the listed files to be executable and run ./setup.sh as an admin

You should see the following message:<br />
**Aurum ascii logo**<br />
Welcome! Let's get everything set up.<br />
Initializing Aurum wallet...<br />
Wallet setup complete.<br />
Okay your wallet's ready.<br />
Run ./client -h or ./client --help to see all commands<br />
Check out the wiki on the github page for more info: https://github.com/SIGBlockchain/aurum_client/wiki<br />

A file called aurum_wallet.json was created. Copy this file and store it in another directory as a backup (not in the aurum directory). __It contains your private key. If you lose you private key, you will lose access to your wallet and it is not recoverable.__<br />

You are now ready to use aurum. If you type ./client -h, you will see the following:<br />
Usage of ./client:<br />
  -info<br />
    	&nbsp;&nbsp;&nbsp;&nbsp;wallet info<br />
  -send string<br />
    	&nbsp;&nbsp;&nbsp;&nbsp;value to send<br />
  -setup<br />
    	&nbsp;&nbsp;&nbsp;&nbsp;set up client<br />
  -to string<br />
    	&nbsp;&nbsp;&nbsp;&nbsp;recipient<br />
  -update<br />
    	&nbsp;&nbsp;&nbsp;&nbsp;update wallet info<br />
  -version<br />
    	&nbsp;&nbsp;&nbsp;&nbsp;client version<br />

* send string means put the address of the person you want to interact with.<br />
* setup is to create a __new wallet__<br />
* update gets the latest transaction data<br />

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
