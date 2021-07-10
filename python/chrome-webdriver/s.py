from selenium import webdriver

options = webdriver.ChromeOptions()
options.add_argument('--headless')
options.add_argument('--no-sandbox')
options.add_argument('--no-gui')

driver = webdriver.Chrome(chrome_options=options)
# driver.get('https://www.meteoblue.com/en/weather/forecast/week/beijing_china_1816670')
driver.get('https://www.baidu.com/')
driver.save_screenshot('screenshot.png')
