Aurum Client
=============
<img src = "assets/aurum_logo_readme.jpg"  alt="drawing" width="200"/>

## Setup
* Navigate to the [releases page](https://github.com/SIGBlockchain/aurum_client/releases)
* Download the archive file compatible with your operating system.
* Extract the archive in a secure folder.
* Run `client --setup` to setup your wallet.
* Your wallet address, balance, and state nonce will be displayed.
* A new file will be generated called `aurum_wallet.json`
* It is strongly recommended you make a copy of your `aurum_wallet.json` and store it in a secure location.
* __This file contains your private key. If you lose your key you will be unable to access your account.__

## Usage
* To display your wallet contents at any time, run `client --info`
* To sync your wallet with the blockchain, run `client --update`
* To send Aurum, run `client --send <AmountToBeSent> --to <RecipientWalletAddress>`

## Project
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
