import pandas as pd

fichero_unificado = 'data/SERVICIO_UNIFICADO_2023.csv'
fichero_parquet = 'data/SERVICIO_UNIFICADO_2023.parquet.gzip'

print("Iniciando el proceso")
#Se lee el archivo en pandas
df = pd.read_csv(fichero_unificado, sep=',', engine='python' , encoding='utf-8', header=0)
print("Leyendo el archivo")

#Se convierte el archivo a parquet
'''table = pa.Table.from_pandas(df)
pq.write_table(table, fichero_parquet)
'''
df.to_parquet(fichero_parquet,
              compression='gzip')  
pd.read_parquet(fichero_parquet) 
print("Finalizado correctamente")
df.head(10)