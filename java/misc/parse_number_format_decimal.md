public String parse_decimal(String value)
{
    while( value.endsWith(",") )
        value = value.substring(0, value.length() - 1);
    while( value.startsWith(",") )
        value = value.substring(1);

    return value.replaceFirst(",", ".").replaceAll(",","");
}