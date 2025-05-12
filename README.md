![Screenshot 2025-05-12 213602](https://github.com/user-attachments/assets/3b50da71-96d9-417f-b7e9-7d3bfb615cfe)




WHERE IS THE KEY?


You can find by youself

here is bonud frida script :


```javascript

Java.perform(function () {
    var CryptoClass = Java.use('com.net2software.encengine.a');

    CryptoClass.$init.overload('java.lang.String', 'java.lang.String').
    implementation = function (key, iv) {
        const maskedKey = "*".repeat(Math.max(0, key.length - 4)) + key.slice(-4);

        const table = `
+------------------+----------------------------------+
| Field            | Value                            |
+------------------+----------------------------------+
| Key              | ${maskedKey.padEnd(32)} |
| IV               | ${iv.padEnd(32)} |
+------------------+----------------------------------+
        `.trim();

        console.log("[*] Creating CryptoClass with:");
        console.log(table);

        return this.$init(key, iv);
    };

    CryptoClass.c.overload('[B').implementation = function (data) {
        console.log("[*] Encrypting data (byte[]):", bytesToString(data));
        return this.c(data);
    };
    CryptoClass.a.overload('java.lang.String').implementation = function (data) {
        console.log("[*] Decrypting data:", data);
        return this.a(data);
    };

    function bytesToString(bytes) {
        let result = "";
        for (let i = 0; i < bytes.length; ++i) {
            result += String.fromCharCode(bytes[i] & 0xff);
        }
        return result;
    }
});


```
