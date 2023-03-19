#!/bin/zsh

now=$(date +"%a %b %d %T %Y")
mkdir "Averil at $now"
cd "Averil at $now"

mkdir -p about_me/personal
mkdir -p about_me/professional
mkdir my_friends
mkdir my_system_info

echo "https://web.facebook.com/averil.primayuda/" > about_me/personal/facebook.txt
echo "https://www.linkedin.com/in/averil-primayuda-a6414b160/" > about_me/professional/linkedin.txt
echo -e "My username: averilp/n$(uname -a)" > my_system_info/about_this_laptop.txt

curl -L -o my_friends/list_of_my_friends.txt "https://gist.githubusercontent.com/tegarimansyah/e91f335753ab2c7fb12815779677e914/raw/94864388379fecee450fde26e3e73bfb2bcda194/list%2520of%2520my%2520friends.txt"
ping -c 1 google.com > my_system_info/internet_connection.txt
