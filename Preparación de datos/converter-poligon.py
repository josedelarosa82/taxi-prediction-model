import pandas as pd
import json
import shapely.geometry
from shapely.geometry import Point, Polygon

fichero_unificado = 'data/SERVICIO_UNIFICADO_2023.csv'
fichero_poligino = 'data/poligonos-localidades.csv'
fichero_parquet = 'data/SERVICIO_UNIFICADO_2023.parquet.gzip'

print("Iniciando el proceso")
#Se lee el archivo en pandas
df = pd.read_csv(fichero_unificado, sep=',', engine='python' , encoding='utf-8', header=0)
print(f"Leyendo el archivo {fichero_unificado}...")

#Cargar el dataset de localidades de bogotá
df_loc = pd.read_csv(fichero_poligino,';', header=0 )
print(f"Leyendo el archivo {fichero_poligino}...")
#Se convierte el archivo a parquet
'''table = pa.Table.from_pandas(df)
pq.write_table(table, fichero_parquet)
'''

#Función para extraer el array interno del json
def convert(param):
  df2 = json.loads(param)
  for val in df2["coordinates"]:
    for val2 in val:
      df2 = val2
  return df2

#Función para descomponer el arreglo interno
def converPolygon(param1):
  np_array = []
  for val in param1:
    np_array+=[(val[0], val[1])]
  return np_array

#Función que buscar un punto dentro de un polígono
def findLocalidad(x, y, array_loc):
  for loc in array_loc:
    polygon = Polygon(loc[5])
    point = Point(x, y)
    #print(polygon.contains(point)) # check if polygon contains point
    #print(point.within(polygon)) # check if a point is in the polygon
    #print(loc[0]," -> ",polygon.contains(point))
    if polygon.contains(point):
      return loc
  return []


df_loc['Geometry'] = df_loc['Geometry'].map(lambda x: convert(x))
df_loc = df_loc.drop(columns=["geo_point_2d"])

df_loc["Polygon"] = df_loc["Geometry"]
df_loc["Polygon"] = df_loc["Polygon"].map(lambda x: converPolygon(x))

x = -74.089700
y =	4.628984
arr_poly = findLocalidad(x,y, df_loc.to_numpy())
if len(arr_poly) > 0:
  print("La localidad es", arr_poly[0])
else:
  print("No hay localidad")

ds_tx = df.to_numpy()
ds_loc = df_loc.to_numpy()

for tx in ds_tx:
    x = tx[6]
    y =	rx[4]
    arr_poly = findLocalidad(x,y, ds_loc)
    if len(arr_poly) > 0:
      print("La localidad es", arr_poly[0])
    else:
      print("No hay localidad")

df.to_parquet(fichero_parquet,
              compression='gzip')  
pd.read_parquet(fichero_parquet) 
print("Finalizado correctamente")