# A list of spotted edge-cases that require special handling

### A. In an OPTIONS section, an entry has more than two coma separated values

Context is a sub-command and a reference to an other sub-command appears in description.

> `manparse` should be able to fetch the corresponding options in the sibling sub-command model.

#### Example (`go-get`)
**man**
```
-a, -n, -v, -x, -p
      The -a, -n, -v, -x, and -p flags have the same meaning as in 'go build'
      and 'go install'.  See go-build(1).
```

**docbook**
``` xml
<varlistentry>
  <term><emphasis role='strong' remap='B'>-a, -n, -v, -x, -p</emphasis></term>
  <listitem>
    <para>The -a, -n, -v, -x, and -p flags have the same meaning as in 'go build'
      and 'go install'.  See <citerefentry><refentrytitle>go-build</refentrytitle><manvolnum>1</manvolnum></citerefentry>.</para>
  </listitem>
</varlistentry>
```

