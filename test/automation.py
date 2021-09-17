
from selenium import webdriver
from selenium.webdriver.common.keys import Keys
import time

browser = webdriver.Chrome()

browser.get("http://127.0.0.1:5500/front-end/")




def addItem():
    print("Adding Test Starting...")
    try:
        time.sleep(2)
        inputArea = browser.find_element_by_xpath("""//*[@id="app"]/p[2]/input""")
        inputArea.send_keys("Modanisa")
        time.sleep(2)

        sendButton = browser.find_element_by_xpath("""//*[@id="passwordHelpInline"]/button""")
        sendButton.click()
        time.sleep(2)
        print("Adding Test Succesful!")
    except NameError:
        print("Something went wrong : "+NameError)

def deleteItem():
    print("Deleting Test Starting...")
    try:
        time.sleep(5)
        willBeDeleted = browser.find_element_by_xpath("""//*[@id="app"]/ul/li[1]/button""")
        willBeDeleted.click()
        time.sleep(2)
        print("Deleting Test Succesful!")
    except NameError:
        print("Something went wrong : "+NameError)

if __name__ == "__main__":
    addItem()
    time.sleep(5)
    deleteItem()


