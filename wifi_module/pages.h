const string BASE_HEADER = "HTTP/1.1 200 OK\r\n"
    + "Content-Type: text/html\r\n\r\n";

const string HTML_START = "<!DOCTYPE HTML>\r\n<html>\r\n";

const string END_TAGS = "</html>\n";

const string REFRESH_HEADER = "<meta http-equiv='refresh' content='1'/>\r\n";

const string NOT_FOUND_PAGE =  "<div>Invalid Request.<br> Try /home.</div>";

string getHomePage() {
    return BASE_HEADER + HTML_START + "<div>test</div>" + END_TAGS;
}

string getNotFoundPage() {
    return BASE_HEADER + HTML_START + NOT_FOUND_PAGE + END_TAGS;
}
