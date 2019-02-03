package resources

const Index = `<html>
<head>
    <title></title>
    <script>
        function addPoint() {
            let e = document.getElementById("points");
            e.insertAdjacentHTML('beforeend', "<p>\n" +
                    "            <label>\n" +
                    "            Point:\n" +
                    "                <input type=\"text\" name=\"point\">\n" +
                    "            </label>\n" +
                    "        </p>");
        }
    </script>
</head>
<body>
<form action="/path" method="post">
    <div id="points">
        <p>
            <label>
            Point:
                <input type="text" name="point">
            </label>
        </p>
        <p>
            <label>
                Point:
                <input type="text" name="point">
            </label>
        </p>
    </div>
    <input type="button" onclick="addPoint()" value="+">
    <input type="submit" value="Get">
</form>
</body>
</html>`
