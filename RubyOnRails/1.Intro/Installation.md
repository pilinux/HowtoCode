## OS: Debian/Ubuntu

### Install ```node```

**Option:** Using ```nvm``` in non-sudo mode

- Link: https://github.com/nvm-sh/nvm
- Restart the terminal
- Available node versions: ```nvm ls-remote```
- Install a specific version of node (example): ```nvm install 12.16.1```


### Install Yarn

- Link 1: https://yarnpkg.com/en/docs/install
- Link 2: https://classic.yarnpkg.com/en/docs/install#debian-stable

```
curl -sS https://dl.yarnpkg.com/debian/pubkey.gpg | sudo apt-key add -
echo "deb https://dl.yarnpkg.com/debian/ stable main" | sudo tee /etc/apt/sources.list.d/yarn.list

sudo apt update
sudo apt install --no-install-recommends yarn
```


### Install Ruby using RVM

- RVM: http://rvm.io/

```
sudo apt install gnupg2
gpg2 --recv-keys 409B6B1796C275462A1703113804BB82D39DC0E3 7D2BAF1CF37B13E2069D6956105BD0E739499BDB
curl -sSL https://get.rvm.io | bash -s stable
sudo usermod -a -G rvm `whoami`
```

- Install ruby requirements: http://rvm.io/rubies/installing

```
rvm autolibs enable
rvm requirements
```

- List known rubies: `rvm list known`

- Install ruby: `rvm install 2.6.5`

- List installed rubies: `rvm list`

- Gem version: `gem -v`

- Gem update: `gem update --system`


### Install Webpack

`gem install webpack`


### Install Rails

- Search available rails: `gem search '^rails$' --all`
- Install rails: `gem install rails -v 6.0.2.2`



---