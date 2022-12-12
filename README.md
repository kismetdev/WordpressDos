# WordpressDos

Exploiting the /xmlrpc.php file on wordpress installations.


This is basically just a rewrite of the [Metasploit module](https://github.com/rapid7/metasploit-framework/blob/3f3bf215600498441701b5d5f4036874f8d3c32d/modules/auxiliary/dos/http/wordpress_xmlrpc_dos.rb) written in Ruby, but in Go for faster speeds (Goroutines ⚡).

## Instructions

```bash
./wpdos -url "https://example.com/xmlrpc.php" -req 10000
```
Parameters:
- `-url` is the Target url (with /xmlrpc.php)
- `-req` is the Amount of Requests (goroutines) to send

Because we are purely using goroutines for each request for the maximum amount of speed, this can get very CPU heavy depending on your specs and the amount of requests.

❗ For security testing purposes only! I am not responsible for any damage done with this tool.❗
## The "Exploit"

The Dos "Exploit" I am using here is not a real exploit per se, but more of a feature that can be exploited.  

The ``XML-RPC API`` is usually used to interact with the Wordpress system (editing posts, etc..).
This can be abused by sending huge amounts of XML requests to the api, which may result in database crashes or even complete server crashes with just a single machine (Dos) on lower end servers. It can also be used to Brute Force a WordPress login.

The xmlrpc.php file should be turned on by default on every Wordpress installation after 3.5, unless handled by something like a .htaccess file or just generally disabled. Having a good [WAF](https://en.wikipedia.org/wiki/Web_application_firewall) will naturally block this DOS attack as well (Rate limit etc..).

