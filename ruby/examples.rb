$global_variable = "this is a global variable"

class HolaMundo

	def initialize()
		@variable_de_instancia = "soy una variable de instancia que solo existe dentro de la clase"
	end

	def saluda()
		puts "Hola Mundo"
		puts $global_variable
		puts @variable_de_instancia
		suma = 2+3
		resta = 2-3
		multiplicacion = 2*3
		division = 2/3 
		exponente = 2**3#dos elevado a la tres
		print "La suma de 2 + 3 = #{suma}, la resta de 2 - 3 = #{resta}"#interpolacion

		cadena = "Hola "
		cadena << " Mundo 2" #concatenacion
		cadena.concat(33)#concatena el caracter que corresponde al nro 33 en ascii
		print cadena #imprime (Hola Mundo 2!)
		print cadena.length #longitud de cadena
		print "abc" <=> "abcdefg" #retorna 1 si la primera cadena es mas larga, -1 si lo es la segunda cadena, 0 si tienen igual longitud (case sensitive)
		print "abc".casecmp("abcdefg")#lo mismo pero no es case sensitive
		print "agus".capitalize #imprime (Agus)
		print "agus".each_char{|caracter| print "_#{caracter}" }#por cada caracter va a la funcion each y utiliza el caracter
		print "MENSAJE PARA CONSOLA".center(40, "-")#----------------MENSAJE PARA CONSOLA----------------
#CONDICIONALES
		if 1 < 15
			puts "1 es menor a 15"
		else
			puts "1 no es mayor a 15"
		end

		if 1 == 1
			puts "1 es igual a 1"
		end

		if not 5 == 2
			puts "5 no es igual a 2"
		end

		if 1==1 and 2==2
			puts "and condition printed"
		end

		if 1==1 or 2==2
			puts "or condition printed"
		end

		unless "a variable" == "Juan"
			print "a variable is not Juan"
		end

		nombre = "Juan"
		if nombre == "Pedro"
			print "es pedro"
			elsif nombre == "Juan"
				print "es juan"

		end

		numero = 115

		respuesta = case numero  
			when 0..100 
				then "esta entre 0 y 100"
			when 101..200 
				then "esta entre 101 y 200"
			when 201..300 
				then "esta entre 201 y 300"
			else
				 "el numero esta fuera del rango 0-300"
		end 
		print respuesta

		numero = 2
		respuesta = case numero
			when 1, 2, 3, 4, 5 then "Entre primeros 5"
			when 6, 7, 8, 9, 10 then "Entre los segundos 5"
		end
		print respuesta
#CICLOS
		for i in(0..10)

			print "#{i},"
			if i == 3
				break
			end
			if i == 0
				next
				#redo
			end

		end

		 puts *(1..10)
		 (1..10).each_char{|i| print i }
		 1.upto(10){|i| print i}
		 10.downto(1){|i| print i}
		 10.times{|i| print i+1}


		 while i<10 do
		 	print i 
		 	i += 1 
		 end

		 begin
		 	print i 
		 	i += 1
		 end while i < 10 #until i>5


		 arreglo = [1, 2, 3]<< 4
		 print arreglo[-1]#4
		 arreglo.push(5)
		 for element if arreglo
		 	puts element
		 end
 
		 arreglo.each do |element|
		 	puts elment
		 end

		 arreglo2 = arreglo.map{|element| element*2 }#new array with modified datas
		 arreglo3 = arreglo.select{|element| element % 2 == 0 }#new array with some filters datas
		 arreglo3 = arreglo.delete_if{|element| element % 2 == 0 }#new array with some filters datas

		 client = 
		 {
		 	"first_name" => "Agus",
		 	"last_name" => "Jimenez",
		 	"age" => 21,
		 	"tel" => 0985246653
		 }

		 client['ruc'] = 12345

		 client.each do |key, value|
		 	puts value
		 end

		 client_keys = client.keys
		 client_values =  client.values
#funciones anonimas
		funcion_anonima = lambda{[param1] "sending #{param1}"}

		funcion_anonima = lambda 
		do [param1] 
			return "sending #{param1}"
		end
		puts funcion_anonima.call("hello world") 
#bloques
		arreglo = [1, 2, 3]
		class Array
			def iterar
				self.each_with_index do |n, i|
					self[i] = yield(n)
				end
			end
		end
		arreglo.iterar do |n|
			n**2 #[1, 4, 9]
		end
#procedure
		
		class Array
			def iterar2(bloque)
				self.each_with_index do |n, i|
					self[i] = bloque.call(n)
				end
			end
		end
		arreglo2 = [4, 5, 6]
		nombre_procedure = Proc.new do |n|
			n*n
		end
		arreglo2.iterar( nombre_procedure )

#metodos

		










	end

	def get_nombre()
		puts "Dame tu nombre"
		nombre = gets
		print "Hola "+nombre
	end
end

#this is a comentario en linea

=begin
	COMENTARIO 
	EN 
	BLOQUE
=end

object = HolaMundo.new()
object.get_nombre
object.saluda