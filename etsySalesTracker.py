"""
PROGRAM TO GET SCARP VIEW COUNT OF ss AND sl
"""
import re
import time
import datetime
import requests
from bs4 import BeautifulSoup
import pandas as pd

scrapeList = [
"https://www.etsy.com/shop/SSweddings",
"https://www.etsy.com/shop/SelineLounge",
"https://www.etsy.com/shop/LeRoseGifts",
"https://www.etsy.com/shop/RoseGoldRebel",
"https://www.etsy.com/shop/ThePaisleyBox",
"https://www.etsy.com/shop/PoppylovePetal",
"https://www.etsy.com/shop/BlushBlossomSugar",
"https://www.etsy.com/shop/ModParty",
"https://www.etsy.com/shop/BlossomSyrup",
"https://www.etsy.com/shop/PrettyRobesShop",
"https://www.etsy.com/shop/LePortDesigns",
"https://www.etsy.com/shop/ShopatBash",
"https://www.etsy.com/shop/ToHappilyEverAfter",
"https://www.etsy.com/shop/ChicPrincessCloset",
"https://www.etsy.com/shop/TotallyBrides",
"https://www.etsy.com/shop/WeddingPartyBee",
"https://www.etsy.com/shop/GlamBridalGifts",
"https://www.etsy.com/shop/MagnoliaBlueSouth",
"https://www.etsy.com/shop/SimplyNameIt"
]

def grabEmBtP(url):
	ct = str(datetime.datetime.now())

	req = requests.get(url)
	soup = BeautifulSoup(req.content,"html5lib")
	soup = soup.prettify()

	sales = re.findall(r'[0-9,]* sales', soup, re.I)[0]
	sales = int("".join(filter(str.isdigit, sales)))
	
	admirers = re.findall(r'[\d,]* admirers', soup, re.I)[0][:-9]

	quality = re.findall(r'quality[\s]*<span[\-\s<>\da-z0-9=\"]*', soup, re.I)[0][-24:-16]
	quality = int("".join(filter(str.isdigit, quality)))
	
	shipping = re.findall(r'shipping[\s]*<span[\-\s<>\da-z0-9=\"]*', soup, re.I)[0][-25:-16]
	shipping = int("".join(filter(str.isdigit, shipping)))

	customerService = re.findall(r'customer\sservice[\s]*<span[\-\s<>\da-z0-9=\"]*', soup, re.I)[0][-25:-16]
	customerService = int("".join(filter(str.isdigit, customerService)))
	
	reviewCount = re.findall(r'reviewCount": "[0-9]*', soup, re.I)[0]
	reviewCount = int("".join(filter(str.isdigit, reviewCount)))


	masterList = [ct[:-6], url[26:], sales, admirers, reviewCount, quality, shipping, 
					customerService]
	return(masterList)


print("       TIME   ,          Shop      ,  sales, admirers, reviewCount, quality, shipping, customerService")


#Main Loop
while True:
	#Infinite Loop begins
	f = open("Etsy.log", "a")

	for item in scrapeList:
		try:

			s = grabEmBtP(item)
			listToStr = ' ,'.join([str(elem) for elem in s]) + "\n"
			print(listToStr)
			f.write(listToStr)
			time.sleep(1)
		except:
			print("An error/exception occured biatch!")
	f.close()


