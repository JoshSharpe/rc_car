using namespace std;

String BASE_HEADER = String("HTTP/1.1 200 OK\r\n") 
    + String("Content-Type: text/html\r\n\r\n");

const String HTML_START = "<!DOCTYPE HTML>\r\n<html>\r\n";

const String END_TAGS = "</html>\n";

const String REFRESH_HEADER = "<meta http-equiv='refresh' content='1'/>\r\n";

const String NOT_FOUND_PAGE =  "<div>Invalid Request.<br> Try /home.</div>";

String getHomePage() {
    return BASE_HEADER + HTML_START + "<body><h2>R/C Car</h2><form action='/movements'>Velocity (between 0 and 1):<br><input type='text' name='velocity' value='0'><br>Turn (between -45 degrees for left and 45 degrees for right):<br><input type='text' name='turn' value='0'><br><br><input type='submit' value='Submit'></form></body>" + END_TAGS;
}

String getNotFoundPage() {
    return BASE_HEADER + HTML_START + NOT_FOUND_PAGE + END_TAGS;
}
