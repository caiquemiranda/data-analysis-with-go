import pandas as pd

df = pd.read_excel('dados.xlsx')
df.to_csv('dados.csv')
