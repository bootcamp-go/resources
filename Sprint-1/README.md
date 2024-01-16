# Instalar MySQL en Mac
Esta guía cubre la instalación de MySQL en Mac, tanto para chips M1/M2 como Intel.

---

## Solución de Alta Prioridad - IntelligentHub
Instalar directamente desde intelligent hub. Parte de las reglas de MeLi

---

## Soluciones Alternativas
### Para Chips Intel
1. **Instalar Homebrew:**
   - Abre la Terminal.
   - Ejecuta `/bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"`.
   - Sigue las instrucciones en pantalla para completar la instalación.
2. **Instalar MySQL:**
   - Una vez instalado Homebrew, ejecuta `brew install mysql`.
3. **Iniciar el Servicio de MySQL:**
   - Ejecuta `brew services start mysql`.
4. **Configurar MySQL:**
   - Ejecuta `mysql_secure_installation`.
   - Sigue las instrucciones para configurar tu instalación de MySQL.

### Para Chips M1/M2 (Apple Silicon)
1. **Descargar MySQL:**
   - Visita la página de descargas de MySQL.
   - Selecciona macOS como sistema operativo.
   - Elige la versión para macOS (ARM64).
2. **Instalar MySQL:**
   - Abre el archivo .dmg descargado y sigue las instrucciones para instalar MySQL.
3. **Agregar MySQL al PATH (si es necesario):**
   - Abre tu archivo `.zshrc` o `.bash_profile` dependiendo de tu shell.
   - Añade `export PATH=/usr/local/mysql/bin:$PATH`.
   - Guarda y cierra el archivo.
   - Ejecuta `source ~/.zshrc` o `source ~/.bash_profile`.
4. **Iniciar MySQL:**
   - Ve a Preferencias del Sistema > MySQL.
   - Inicia el servidor de MySQL.
5. **Configurar MySQL:**
   - Abre la Terminal.
   - Ejecuta `mysql_secure_installation`.
   - Sigue las instrucciones para configurar tu instalación de MySQL.

---

## Verificar la Instalación
- Ejecuta `mysql -u root -p`.
- Introduce tu contraseña cuando se te solicite.
- Deberías tener acceso al prompt de MySQL.

---

# Inicializar la DB
En esta guía seguiremos los pasos para inicializar la DB del proyecto.

## Crear la DB y el Usuario
- Ejecuta `./db_init.sh`. Esto creará la DB y el usuario con los permisos necesarios para el proyecto.

## Cargar los Datos (opcional)
- Ejecuta `./db_load.sh`. Esto resetea los datos y cargará algunos nuevos de prueba en la DB.