# Protobuf

Este proyecto de ejemplo muestra cómo utilizar **protobuf** en Go para serializar y deserializar datos.

## Descripción

**Protobuf** (Protocol Buffers) es un mecanismo de serialización de datos desarrollado por Google. Proporciona una forma eficiente y flexible de intercambiar datos estructurados entre diferentes sistemas y lenguajes de programación. Este proyecto de ejemplo demuestra cómo utilizar **protobuf** en Go para definir mensajes y generar código para serializar y deserializar esos mensajes.

## Instalación

1. Asegúrate de tener **Go** instalado en tu sistema. Puedes obtenerlo en **https://golang.org/**.

2. Clona este repositorio en tu máquina local.

    ```bash
    git clone https://github.com/mendoza-j8p/protobuf.git
    ```

3. Ve al directorio del proyecto.

    ```bash
    cd protobuf
    ```

4. Instala las dependencias del proyecto.

    ```bash
    go mod download
    ```

## Creado

1. Define tus mensajes protobuf en el archivo **person.proto** siguiendo la sintaxis de **protobuf**.

2. Ejecuta el siguiente comando:

    ```bash
    go get github.com/golang/protobuf/proto
    ```

    Después de ejecutar **go get**, Go descargará el código fuente del paquete desde el repositorio correspondiente en GitHub y lo almacenará en el directorio GOPATH de tu sistema. A partir de ahí, podrás importar y utilizar el paquete protobuf en tus programas Go.

Si te encuentras con errores relacionados con el acceso a comandos o paquetes en tu sistema, es probable que necesites agregar ciertas rutas al PATH de tu sistema.
Por ejemplo, si estás trabajando con Go y te falta el paquete "protobuf", puedes seguir estos pasos:

- Ejecuta el siguiente comando para agregar la ruta al PATH temporalmente:

  * En sistemas Unix/Linux:

  ```bash
  export PATH=$PATH:/ruta/a/protobuf
  ```

  * En sistemas Windows:

  ```bash
  set PATH=%PATH%;C:\ruta\a\protobuf
  ```

    Reemplaza "/ruta/a/protobuf" o "C:\ruta\a\protobuf" con la ubicación real del paquete o comando que falta.

Como opción alternativa, puedes agregar la ruta de forma permanente al PATH de tu sistema:

- En sistemas Unix/Linux:

    Abre el archivo **~/.bashrc** o **~/.bash_profile** en un editor de texto y agrega la siguiente línea al final del archivo:

    ```bash
    export PATH=$PATH:/ruta/a/protobuf
    ```

    Reemplaza "/ruta/a/protobuf" o "C:\ruta\a\protobuf" con la ubicación real del paquete o comando que falta.

    Después de realizar los cambios en el archivo ~/.bashrc o ~/.bash_profile, es necesario recargar la configuración para que los cambios tengan efecto en la sesión actual de la terminal.

    Para esto ejecuta el comando source seguido del nombre del archivo:

    ```bash
    source ~/.bashrc
    ```

    o

    ```bash
    source ~/.bash_profile
    ```

    De esta manera, se recargará la configuración del archivo y el sistema utilizará la nueva configuración del PATH que has agregado.

- En sistemas Windows:

    Haz clic derecho en "Este equipo" o "Mi PC" y selecciona "Propiedades". Luego, ve a "Configuración avanzada del sistema" y haz clic en "Variables de entorno". En la sección "Variables del sistema", busca la variable "Path", haz clic en "Editar" y agrega la ruta al final de los valores existentes, asegurándote de separarla con un punto y coma (;) del valor anterior.

    Guarda los cambios y cierra la ventana de Propiedades del sistema.

3. Genera el código Go correspondiente ejecutando el siguiente comando:

    ```bash
    protoc --go_out=. person.proto
    ```

    Esto generará el archivo **person.pb.go** en el directorio actual.
    - **protoc**: es el comando para ejecutar el compilador de Protocol Buffers.
    - **--go_out**: es una opción que se le pasa al compilador para indicar que queremos generar código Go.
    - **.**: es el argumento que se le pasa a la opción **--go_out** y especifica el directorio de salida donde se generará el archivo Go.
    - **person.proto**: es el archivo **.proto** que contiene la definición de los mensajes protobuf.

4. En el archivo **main.go**, puedes importar el paquete generado **github.com/mendoza-j8p/protobuf** y utilizar los mensajes definidos en **person.pb.go** para serializar y deserializar datos.

5. Ejecuta el programa principal.

    ```bash
    go run main.go
    ```

    Esto mostrará un ejemplo de uso de la serialización y deserialización de los mensajes **protobuf**.

## Contribuciones

¡Las contribuciones son bienvenidas! Si deseas mejorar este proyecto, crea un pull request con tus cambios.
