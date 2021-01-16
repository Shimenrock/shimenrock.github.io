---
title: "What Is SQL Injection and How to Stop It"
published: false
related: true
toc: true
toc_sticky: true
categories: 
  - Injection
---  

Data breaches are a common occurrence. As described in The Securing Account Details story, SQL injection is a simple way to access data from exposed sites.

数据泄露是常见的情况。 如“保护帐户详细信息”故事中所述，SQL注入是一种从暴露的站点访问数据的简单方法。

How easy is SQL injection and how bad can it be?

SQL注入有多容易，有多糟糕？

In this post we'll take a look at how it's possible. We'll see how easy it is to access information from a database that is vulnerable to SQL injection. We'll finish up by showing how you can prevent it.

在这篇文章中，我们将探讨它的可能性。 我们将看到从容易受到SQL注入攻击的数据库访问信息有多么容易。 我们将向您展示如何预防这种情况。

Let's start with a simple HTML form. It accepts an email address. It passes this to a Java servlet. This looks up subscription details for the email.

让我们从一个简单的HTML表单开始。 它接受一个电子邮件地址。 它将其传递给Java servlet。 这将查找电子邮件的订阅详细信息。

The form is:

```
<body>
  <form action='HelloInjection'>

  <input type='text' name='emailAddress' value=''/>
  <input type='submit' value='Submit'/>
</form>
</body>
```

The following Java code processes the values:
```
protected void doGet(HttpServletRequest request, HttpServletResponse response) throws ServletException, IOException {
   PrintWriter out = response.getWriter();
   Connection conn = null;

   try {
      Class.forName ("oracle.jdbc.driver.OracleDriver");
      String url = "jdbc:oracle:thin:@//db_server:1521/service_name";
      String userName = "XXXX";
      String password = "XXXX";

      conn = DriverManager.getConnection(url, userName, password);

      String email = request.getParameter("emailAddress");
      Statement st = conn.createStatement();
      
      String query= "select * from email_subscriptions where email_address = '" + email + "'";

      out.println("Query : " + query + "<br/>");

      ResultSet res = st.executeQuery(query);
 
      out.println("Results:" + "<br/>");
      while (res.next()) {
     String s = res.getString("email_address");
   out.println(s + "<br/>");
      }

   } catch (Exception e) {
      e.printStackTrace();
   } 
}
```

The key part of this the query:
```
 String query= "select * from email_subscriptions where email_address = '" + email + "'";
```

This takes whatever the user inputs and concatenates it to the string. For example, if we submit chris.saxon@oracle.com, the query becomes:
 select * from email_subscriptions where email_address = 'chris.saxon@oracle.com'
What's so bad about this?

We can supply whatever we want to go between the quotes. This enables us to add more conditions to the where clause. To do this, just close the quotes. Then add our additional criteria. For example, if we submit the following:
```
 ' or 1='1
 ```
This query becomes:
```
 select * from email_subscriptions where email_address = '' or 1='1'
```

The expression 1='1' is always true. True or anything is true. Therefore this query returns all the email addresses stored in the table! The code loops through all the results, displaying them all.

表达式1 ='1'始终为true。 正确或任何事情都是正确的。 因此，此查询返回表中存储的所有电子邮件地址！ 代码循环遍历所有结果，显示所有结果。

This means anyone can get their hands on the table's contents. If it contains sensitive or private data we're in trouble.

这意味着任何人都可以使用表格中的内容。 如果其中包含敏感数据或私人数据

But, you may say, our application only uses string concatenation for queries against public or non-sensitive data. For example, lookup tables such as countries and currencies. It can't be that bad, can it?

但是，您可能会说，我们的应用程序仅使用字符串连接来查询公共数据或非敏感数据。 例如，查询表，例如国家和货币。 不会那么糟糕，是吗？

Yes it can.

是的，它可以

Remember that you can union queries together. With this simple operation, we can view the contents of any table in the database user has access to!

是的，它可以

What if they don't know the names of our tables?

如果他们不知道我们表的名称怎么办？

That's not a problem. Just submit the following:
```
 ' union all select table_name from all_tables where 1 = '1
```

And the query becomes (formatted for clarity):
```
 select * from email_subscriptions where email_address = '' 
        union all 
        select table_name from all_tables where 1 = '1
``
Hmmm. This returns all the tables the user has access to. Using a similar query, we can also find all the columns they can access. Armed with this information a hacker could union the original query with any other table, potentially leaving your whole database exposed.

这将返回用户有权访问的所有表。 使用类似的查询，我们还可以找到它们可以访问的所有列。 有了这些信息，黑客就可以将原始查询与其他任何表合并，从而可能使整个数据库暴露在外。

So how can we stop this?

那么我们怎样才能阻止这种情况呢？

Simple. Use bind variables.

简单。 使用绑定变量。

To do this in the Java code above, change the Statement to a PreparedStatement. Then modify the query string to have a question mark (?) instead of concatenating the email address into it. Finally set the value of the variable to the supplied email address.

要在上面的Java代码中执行此操作，请将Statement更改为PreparedStatement。 然后将查询字符串修改为带有问号（？），而不是将电子邮件地址串联到其中。 最后，将变量的值设置为提供的电子邮件地址。

Putting this all together gives:
```
    String query= "select * from email_subscriptions where email_address = ?";
    PreparedStatement st = conn.prepareStatement(query);
```
 st.setString(1, email);
 ResultSet res = st.executeQuery();
```

Now, regardless of what someone types in the form, the query will always be:

现在，无论有人在表单中键入什么，查询将始终是

```
 select * from email_subscriptions where email_address = ?
```

Hacking attempts such as ' or 1='1 now return no results.

Going further ensure your application user and table owner are different database users. Then follow the principle of least privilege to ensure that application users only have access to the tables they need. This limits what hackers can see if they do manage to find a loophole.

进一步确保您的应用程序用户和表所有者是不同的数据库用户。 然后遵循最小特权原则，以确保应用程序用户只能访问他们需要的表。 这限制了黑客在设法找到漏洞时可以看到的内容。

Even better, build PL/SQL APIs to fetch and modify data. Grant application users permissions to just these APIs. With no direct access to tables, you've further decreased the risk of people accessing something they shouldn't.

更好的是，构建PL / SQL API来获取和修改数据。 授予应用程序用户仅这些API的权限。 由于不能直接访问表，因此可以进一步降低人们访问不该访问的内容的风险。

Above all: to ensure your data is safe, use bind variables!

最重要的是：为确保数据安全，请使用绑定变量！

