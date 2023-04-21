import pandas as pd

fichero_unificado = 'data/SERVICIO_UNIFICADO_2023.csv'
fichero_parquet = 'data/SERVICIO_UNIFICADO_2023.parquet.gzip'

print("Iniciando el proceso")
#Se lee el archivo en pandas
df = pd.read_csv(fichero_unificado, sep=',', engine='python' , encoding='utf-8', header=0)
print(f"Leyendo el archivo {fichero_unificado}...")

df.to_parquet(fichero_parquet,
              compression='gzip')  
print("Finalizado correctamente.")