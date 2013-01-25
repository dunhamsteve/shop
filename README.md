# Shop

A simple go program to update a "ShopShop" shopping list.

ShopShop is a free iOS shopping list program.  It stores its data in a binary plist format 
within a DropBox folder. This app allows you to manage one of your lists via the command line.

To use, you should have a `$HOME/Dropbox/ShopShop/Shopping List.shopshop` file, created
by ShopShop.

## Usage

    $ shop

Will list current items.

    $ shop add foobar

Will add foobar to your shopping list.

    $ shop rm 3

Will remove the item at index 3.

    $ shop co

Will remove items marked as done.
