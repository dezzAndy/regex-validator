using System.Runtime.InteropServices;

namespace ValidadorGUI;

public static class NativeValidator
{
#if WINDOWS
    private const string LibName = "validador";
#else
    private const string LibName = "libvalidador";
#endif

    [DllImport(LibName, CallingConvention = CallingConvention.Cdecl)]
    private static extern bool ValidarTelefono(string input);

    [DllImport(LibName, CallingConvention = CallingConvention.Cdecl)]
    private static extern bool ValidarCorreo(string input);

    [DllImport(LibName, CallingConvention = CallingConvention.Cdecl)]
    private static extern bool ValidarCurp(string input);

    [DllImport(LibName, CallingConvention = CallingConvention.Cdecl)]
    private static extern bool ValidarPass(string input);

    [DllImport(LibName, CallingConvention = CallingConvention.Cdecl)]
    private static extern bool ValidarRfc(string input);

    [DllImport(LibName, CallingConvention = CallingConvention.Cdecl)]
    private static extern bool ValidarIP(string input);

    [DllImport(LibName, CallingConvention = CallingConvention.Cdecl)]
    private static extern bool ValidarCumpleanos(string input);

    [DllImport(LibName, CallingConvention = CallingConvention.Cdecl)]
    private static extern bool ValidarPlaca(string input);

    public static bool Validar(TipoCadena tipo, string input) => tipo switch
    {
        TipoCadena.Telefono => ValidarTelefono(input),
        TipoCadena.Correo => ValidarCorreo(input),
        TipoCadena.Curp => ValidarCurp(input),
        TipoCadena.Password => ValidarPass(input),
        TipoCadena.Rfc => ValidarRfc(input),
        TipoCadena.Ip => ValidarIP(input),
        TipoCadena.Cumpleanos => ValidarCumpleanos(input),
        TipoCadena.Placa => ValidarPlaca(input),
        _ => false,
    };
}