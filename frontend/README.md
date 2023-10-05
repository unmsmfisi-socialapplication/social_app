# App Social
## Next.js App Router
Next.js para manejar la navegación en una aplicación web. Next.js proporciona una forma sencilla y poderosa de crear rutas y navegar entre páginas en una aplicación React.

## Estructura de Carpeta

  ```
├── app
│   ├── page.tsx
│   └── test
├── components
│   ├── atoms
│   ├── index.tsx
│   ├── molecule
│   └── organismos
├── data
│   ├── api
│   └── entities
├── domain
│   ├── repositories
│   └── usecases
└── utilities
    ├── Contant.tsx
    ├── Functions.tsx
    └── Validation.tsx
 ```
### APP
En la carpeta app se encuentra el archivo principal page.tsx y tambien las carpetas deredireccion en Next este  sistema lo encuentras como App Router 
### Components

En este folder irán todos los componentes reutilizable del proyecto ,  cada componente esta desarrollado para ser utilizado en diferentes paginas.Por ello se trabajara sobre nomenclatura atomic desing el cual consta 
de atomos , moleculas , organismos las molecualas o organismos se colcoara en el index.tsx como exportables
  ```

| --componentes
    | -- atoms
       | -- Button
          | -- Button.js
          | -- Button.scss
   | -- moléculas
       | -- Card
          | -- Card.js
          | -- Card.styles.js
   | -- organismos
       | -- Header
          | -- Header.js
          | -- Header.scss
       | -- Footer
          | -- Footer.js
          | -- Footer.scss
  ```
### Data
Se encuentra las apis , estos metodos puede ser get , delete , put y post para recibir los servicios , entities se colocara las interfaces sera el cuerpo de los productos que manejaremos esta parte es inmutable por ello hay adaptadores por ejemplo create'Error' se adaptara la entidad dependiendo sea necesario pero no la interfaz



### Domain
Se juntaron los servicios especificos por la logica de negocio aca se estrucutrizara la data junto con las interfaces de las entidades y se documentatara estos para ser compartida en todo el proyecto en repositories
### utilities
Se desacoplara funciones que permiten por ejemplo valdiacion , transformacion y seran llamadas en diferentes partes del proeycto con el objetivo de minimizar la mayor cantidad de codigo , asu vez las constantes sera informacion parametrizada reusada varias veces en el proyecto como fechas , nombre de los productos y de esta forma proteger la informacion del flujo y se pueda cambiar en el futuro en el caso la logica de negocio lo necesite



## Creación de Páginas 
Next.js maneja las rutas mediante la carpeta app luego se procedera a crear tu carpeta y colocar uyn archivo page.tsx . Cada archivo JavaScript (o TypeScript) en esta carpeta se convierte automáticamente en una página accesible mediante una URL específica. Por ejemplo, debe existir una carpeta test y dentro un archivo llamado page.tsx lo que creara la página /test/page.tsx.
  ```
  src
  |_app
  | |_test
  | | |   page.tsx
  | | home
  | | |_admin
  | | | |   page.tsx
  ```

  ```js
    export default function testPage() {
    return <h1>Hello, Next.js!</h1>
  }
  ```
Es necesario usar 'use client' para realizar accion y hooks 
  ```js
  'use client'
import { useState } from "react"
export default function TestPage() {
    const [count, setCount] = useState(0);

    const handleCount = () => {
        setCount(count + 1);
        alert(count);
    }
    return (
    <main>
        <h1>Test Page</h1>
        <button onClick={handleCount}>presioname</button>
    </main>
  )
}
  ```

## Creación de template y usarlo en paginas 
El archivo layout.tsx en un template en el cual pdoemos agregar pie de pagina y encabezado como tambien lo podemos modificar  para ello se crear difrentes templates dependiendo lo que se necesita y estos sera usado en las paginas  
  ```js
export default function RootLayout({
    children,
  }: {
    children: React.ReactNode
  }) {
    return (
      <html lang="en">
        <body>{children}</body>
      </html>
    )
  }
  ```
  el archivo hijo childdren signifca que podremos colocar pagians interiormente
```js
'use client'
import { useState } from "react"
import Layout from "../layout"
export default function TestPage() {
    const [count, setCount] = useState(0);

    const handleCount = () => {
        setCount(count + 1);
        alert(count);
    }
    return (
        <Layout>
            <h1>Test Page</h1>
            <button onClick={handleCount}>presioname</button>
        </Layout>
      );
}
```
## Como funciona
El enrutador de Next.js simplifica la navegación en tu aplicación al proporcionar una estructura de ruta clara y fácil de usar. Aquí hay algunos conceptos clave:



## Herramientas del proyecto.

| Herramienta  | Descripción                                  |
| ------------ | -------------------------------------------- |
| npm 16       | Para instalar todos los módulos del proyecto |
| BEM          | Block Element Modifier        |
| scss         | Compilador de las hojas de estilo            |
| vscode       | IDE para desarrollo                          |

## Estilo BEM
## Metodología BEM (Block Element Modifier)

La Metodología BEM (Block Element Modifier) es una convención de nomenclatura de clases CSS 

- **Bloque Principal**: El bloque principal es el componente o elemento de nivel superior que estamos estilizando. Por ejemplo, en nuestro proyecto, `.card` podría ser un bloque principal.

- **Elemento**: Los elementos son partes más pequeñas que forman parte del bloque principal. Usamos dos guiones bajos (`__`) para separar el bloque del elemento. Por ejemplo, `.card__title` es un elemento dentro del bloque `.card`, y `.card__button` es otro elemento dentro del mismo bloque.

- **Modificador**: Los modificadores se utilizan para aplicar variaciones o modificaciones a los elementos o bloques. Usamos dos guiones medios (`--`) para separar el elemento o bloque del modificador. Por ejemplo, `.card__button--primary` es una modificación del elemento `.card__button` que le da un estilo primario.

Utilizamos BEM para mantener nuestros estilos CSS organizados y evitar conflictos. Al seguir esta convención, nuestros estilos son más predecibles y fáciles de mantener a medida que desarrollamos y expandimos nuestro proyecto.

Si deseas aprender más sobre BEM, consulta la [documentación oficial de BEM](https://en.bem.info/methodology/).


```
.bloque__elemento--modificador
```
```scss
.bloque{
    margin:0 auto;

    &__boton{
        border: 1px solid black;
        &--rojo{
            background:red;
        }    
    }
}
```


## Estilo de código

### React

- Los imports de librerías de terceros van primero, luego se deja una línea de espacio para comenzar importar los componentes o librerías del proyecto. De preferencia agrupar los imports por tipo: imports relacionado a redux, a componentes, a librerias, etc.
- Los componentes son exportados como funciones:
  ```js
  export default function ComponentName() {
    return <div>Componente</div>;
  }
  ```
- Los Hooks van al comienzo del componente, y son agrupados por tipo: los hooks de libreria y por último hooks de estado.

  ```js
  export default function ComponentName() {
    const navigator = useNavigation();
    const [open, setOpem] = useState(true);
    const [loading, setLoading] = useState(false);
  }
  ```

- Luego de los hooks, siguen las funciones para la aplicación:

  ```js
  export default function ComponentName() {
    const dispatch = useDispatch();
    const navigator = useNavigation();

    const clickHandler = () => {
      console.log(click);
    };
  }
  ```

- Por último viene los useEffect:

  ```js
  export default function ComponentName() {
    const dispatch = useDispatch();
    const navigator = useNavigation();

    const clickHandler = () => {
      console.log(click);
    };

    useEffect(() => {
      dispatch(getUser());
    }, []);
  }
  ```




## Available Scripts

In the project directory, you can run:

### `npm run dev`

Runs the app in the development mode.<br />
Open [http://localhost:3000](http://localhost:3000) to view it in the browser.

The page will reload if you make edits.<br />
You will also see any lint errors in the console.
