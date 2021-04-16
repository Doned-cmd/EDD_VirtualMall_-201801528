export class Usuario {
    Dpi:number
    Password:string
}
export class UsuarioLogeado {
    Dpi:number
    Nombre:string
    Correo: string
    Password:string
    Cuenta: string
}
export class MostrarUsuario {
   Buleano: boolean
   Usuario: UsuarioLogeado
}