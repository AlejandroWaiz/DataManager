# DataManager

CloudFunction that will store data when a excel file is saved into a predefined google cloud storage bucket.

Me acabo de dar cuenta que llamé como orquestador a una pieza que no orquesta, sino que funciona mas como nucleo central
Creo que en vez de llamarlo orquestador deberia ser quien reciba y despache. 

=>

Solucionado el orquestador: ahora se encarga de recibir un data_reader y un data_saver y maneja que cada uno cumpla sus funciones. Actualmente el obtener el excel desde Cloud Storage está en la parte del orquestador lo cual podria sacarse
de ahi y delegarselo al data_reader para crear modulos mas independientes, y eventualmente podria separar la logica 
de cloud storage a un paquete externo pero por ahora esto funciona y sinceramente tengo sueño, son las 2:22 am xD solo falta
hacer las pruebas correspondientes

