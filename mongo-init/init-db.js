// Conectar a DB
db = db.getSiblingDB('taskdb');

// Insertar tareas de ejemplo
// *Generadas artificialmente
db.tasks.insertMany([
  {
    title: "Configurar Docker",
    description: "Finalizar el despliegue con Docker Compose",
    completed: true,
    tags: ["docker", "devops"],
    createdAt: new Date()
  },
  {
    title: "Probar el Frontend",
    description: "Verificar que la paleta de colores se vea correctamente",
    completed: false,
    tags: ["nuxt", "ui"],
    createdAt: new Date()
  },
  {
    title: "Aprender Go",
    description: "Estudiar punteros y estructuras",
    completed: false,
    tags: ["backend", "go"],
    createdAt: new Date()
  },
  {
    title: "Implementar Autenticación JWT",
    description: "Crear el middleware en Gin para proteger las rutas de la API usando JSON Web Tokens.",
    completed: false,
    tags: ["backend", "seguridad", "go"],
    createdAt: new Date()
  },
  {
    title: "Escribir pruebas unitarias",
    description: "Añadir tests en Go para la capa de servicios (TaskService) asegurando que la paginación calcule bien los saltos.",
    completed: false,
    tags: ["backend", "testing"],
    createdAt: new Date()
  },
  {
    title: "Configurar GitHub Actions",
    description: "Crear un pipeline de CI/CD que ejecute los tests y valide el formato del código cada vez que se haga push a la rama principal.",
    completed: false,
    tags: ["devops", "ci-cd"],
    createdAt: new Date()
  },
  {
    title: "Optimizar consultas MongoDB",
    description: "Crear índices en los campos 'completed' y 'tags' para mejorar el rendimiento de las consultas filtradas.",
    completed: false,
    tags: ["database", "mongodb"],
    createdAt: new Date()
  },
  {
    title: "Diseñar la UI de tareas",
    description: "Crear un diseño atractivo para la lista de tareas, incluyendo colores y tipografías que mejoren la experiencia del usuario.",
    completed: false,
    tags: ["frontend", "diseño"],
    createdAt: new Date()
  },
  {
    title: "Implementar búsqueda de tareas",
    description: "Añadir una barra de búsqueda en el frontend que permita filtrar tareas por título o descripción utilizando la API de búsqueda.",
    completed: false,
    tags: ["frontend", "backend", "busqueda"],
    createdAt: new Date()
  },
  {
    title: "Desplegar en Microsoft Azure",
    description: "Implementar el despliegue de la aplicación en la nube de Microsoft Azure.",
    completed: false,
    tags: ["devops", "azure"],
    createdAt: new Date()
  },
  {
    title: "Comprar café",
    description: "Recargar suministros. Fundamental para que el servidor (y el desarrollador) sigan funcionando.",
    completed: false,
    tags: ["urgente", "personal"],
    createdAt: new Date()
  }
]);

print('Base de datos inicializada con éxito');