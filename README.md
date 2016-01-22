# Civo Command-Line Client

This utility is for interacting with the Civo Cloud API provided by Absolute DevOps Ltd. In order to use the API you need an API token, which should have been sent to you by your provider. If you don't have this, please contact them.

## Installation/overview

The first step should be to download the client.  The easiest way at the moment is to visit the [GitHub Releases](https://github.com/absolutedevops/civo/releases) and download the latest version as a Zip file.  This will include folders called `macosx` and `linux` containing the appropriate binaries.  You will need to copy the correct one to a place on your path, `/usr/local/bin` is often a good choice.

You'll then need to run a command in a Terminal to register your token. Let's take the example that your company is called "Acme Widgets" and your token is "123456789012345678901234567890". You need to give the token a short reference when saving it, such as `acme`:

```
civo tokens save -n acme -k 123456789012345678901234567890
```

You will then need to set this as your default token to be used in all future requests with:

```
civo tokens default -n acme
```

Now you are free to use the remaining commands in the system. We'll work through the most common ones below, the rest are normally used by Civo administrators (and the permission levels associated with your token in CIvo won't allow you to make them).

In order to discover the available commands you can use `civo -h` to list the available commands, then use `civo [command] -h` to list sub commands and so on further down the line.  For example:

```
civo -h
civo instances -h
civo instances create -h
```

In this way all the possible things you can do with the client are discoverable.


## SSH keys

One of the first things you'll likely want to do is upload your SSH public key, so that you can SSH in to new instances - you can't create a new instance without this step.

Assuming your public key is in `~/.ssh/id_rsa.pub` (if it isn't, you'll probably know why and where it is) you can upload this with:

```
civo sshkey upload --name default --public-key ~/.ssh/id_rsa.pub
```

If you want to remove a public key (say you are replacing it with a new one), you can do this with:

```
civo sshkey delete --name default
```

**Note:** This won't remove it from your currently running instances, it will only affect new instances created.


## Choosing the specification of an instance
When creating an instance, you'll need to specify items such as the size of instance, which region to create it in (if your provider supports multiple regions) and the template to use (from the available operating systems, versions and layered applications).

The information on all of these are available by running `civo size`, `civo region` and `civo template`.  The output of the command will give you the key to use when creating the instance.  For example (and these are subject to change):

```
$ civo size
+-----------+----------------------------------------------------+
|   Name    |                   Specification                    |
+-----------+----------------------------------------------------+
| g1.xsmall | Extra Small - 512MB RAM, 1 CPU Core, 20GB SSD Disk |
| g1.small  | Small - 1GB RAM, 2 CPU Cores, 50GB SSD Disk        |
| g1.medium | Medium - 2GB RAM, 4 CPU Cores, 100GB SSD Disk      |
| g1.large  | Large - 4GB RAM, 6 CPU Cores, 150GB SSD Disk       |
| g1.xlarge | Extra Large - 8GB RAM, 8 CPU Cores, 200GB SSD Disk |
+-----------+----------------------------------------------------+

$ civo template
+--------------------+------------------------------------------------------------------------+
|         ID         |                              Description                               |
+--------------------+------------------------------------------------------------------------+
| centos-7           | CentOS version 7 (RHEL open source clone)                              |
| ubuntu-14.04-vesta | Canonical's Ubuntu 14.04 with the Vesta Control Panel                  |
| ubuntu-14.04       | Canonical's Ubuntu 14.04 installed in a minimal configuration          |
+--------------------+------------------------------------------------------------------------+
```

