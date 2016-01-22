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


## Managing instances

To view the list of your currently running instances you can simply run:

```
civo instance
```

This will output a table listing the instances currently in your account:

```
+----------+-------------------+----------+-------------------------------+--------+------+--------------+
|    ID    |       Name        |   Size   |         IP Addresses          | Status | User |   Password   |
+----------+-------------------+----------+-------------------------------+--------+------+--------------+
| 8043d0e7 | test1.example.com | g1.small | 10.0.0.2=>31.28.88.103        | ACTIVE | civo | jioAQfSDffFS |
+----------+-------------------+----------+-------------------------------+--------+------+--------------+
```

Creating an instance is a simple command away (remember, if you can't remember the parameters `civo instance create -h` is there to help you) using something like:

```
civo instance create --name test2.example.com --size g1.small \
  --region svg1 --ssh-key default --template ubuntu-14.04 --public-ip
```

If you don't specify a name, a random one will be created for you.

**Note:** Specifying the name will set the hostname on the machine but won't affect DNS resolution, currently that's up to you to provide separately.

If you decide you don't need an instance any more you can remove it by simply calling `civo instance destroy` passing in either the ID or the name, using the details above as an example:

```
civo instance destroy 8043d0e7
civo instance destroy test1.example.com
```

**Note:** The machine will be forever destroyed at this point, you can't get the data back from the hard drive afterwards.

If your machine gets stuck you can restart it with (again using either the ID or the name):

```
civo instance reboot 8043d0e7
```

If it's *really* stuck (i.e. hard kernel lock) then you can do the cloud equivalent of unplugging it and plugging it back in with the addition of the hard switch:

```
civo instance reboot --hard 8043d0e7
```


## Quota

All Civo users have a limited quota applied to their account (to stop errant scripts from filling up the cloud with a million instances).  You can view your current quota using a command like this:

```
$ civo quota
+---------------------------------------+------+-------+
|                 Title                 | Used | Limit |
+---------------------------------------+------+-------+
| Number of instances                   |    0 |    10 |
| Total CPU cores                       |    0 |    20 |
| Total RAM (MB)                        |    0 |  5120 |
| Total disk space (GB)                 |    0 |   250 |
| Disk volumes                          |    0 |    10 |
| Disk snapshots                        |    0 |    30 |
| Public IP addresses                   |    0 |    10 |
| Private subnets                       |    0 |     1 |
| Private networks                      |    0 |     1 |
| Security groups                       |    0 |    10 |
| Security group rules                  |    0 |   100 |
| Number of ports (network connections) |    0 |    20 |
+---------------------------------------+------+-------+
```

If you want to increase them, contact your Civo cloud provider.