import time 

x = 'abcf'
# if elif  else
if x == 'abc' :
    print ('成功')
else:
    print ('no')


chinese_zodiac = '猴鸡狗猪鼠牛虎兔龙蛇马羊'
# year = int(input('输入年份:'))
# print(chinese_zodiac[year % 12])

for cz in chinese_zodiac:
    print(cz)

for i in range(1,13):
    print(i)

for year in range(2000,2019):
    print ('%s 年的生肖是 %s' %(year,chinese_zodiac[year % 12]))

# break continue 跳过当前循环
num = 5 
while True:
    num = num + 1
    if num == 10:
        continue
    print(num)
    time.sleep(1)
