# wordpress

Simple wordpress template with PHP 7 and MYSQL extension enabled.
Just make sure to have your entire wordpress files in the current
directory, then just type 'openode deploy'.

Make sure to set these variables in your wp-config.php:

define( 'WP_HOME', 'http://your-hostname/' );
define( 'WP_SITEURL', 'http://your-hostname/' );

We recommend using an external MYSQL database.
Example free Mysql DB: https://www.freemysqlhosting.net/
