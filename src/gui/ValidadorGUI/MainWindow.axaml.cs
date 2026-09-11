using System;
using Avalonia.Controls;
using Avalonia.Interactivity;
using Avalonia.Media;

namespace ValidadorGUI;

public partial class MainWindow : Window
{
    private static readonly TipoCadena[] Tipos =
    [
        TipoCadena.Telefono,
        TipoCadena.Correo,
        TipoCadena.Curp,
        TipoCadena.Password,
        TipoCadena.Rfc,
        TipoCadena.Ip,
        TipoCadena.Cumpleanos,
        TipoCadena.Placa,
    ];

    public MainWindow()
    {
        InitializeComponent();
        ValidarButton.Click += ValidarButton_Click;
    }

    private void ValidarButton_Click(object? sender, RoutedEventArgs e)
    {
        string entrada = InputBox?.Text ?? string.Empty;
        var tipo = Tipos[Math.Clamp(TipoBox?.SelectedIndex ?? 0, 0, Tipos.Length - 1)];

        bool valido = NativeValidator.Validar(tipo, entrada);

        ResultText!.Text = valido ? "✔ Válido" : "✘ Inválido";
        ResultText.Foreground = new SolidColorBrush(valido
            ? Color.Parse("#2E7D32")
            : Color.Parse("#C62828"));
    }
}