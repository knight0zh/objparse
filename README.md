# objparse
objparse 是一个简易的obj文件解析器。把obj文件转换成二维地图数组。

### Unity World

![avatar](./snapshots/maze.jpg)

### Objparse print World

```
        XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX
        XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX
        XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX
        XXX..............................................XXXXXXXXX..................XXXXXXXXX.............XX
        XXX..............................................XXXXXXXXX..................XXXXXXXXXX............XX
        XXX..............................................XXXXXXXXX.................XXXXXXXXXXXX...........XX
        XXX..............................................XXXXXXXXX................XXXXXXXXXXXXXX..........XX
        XXX..............................................XXXXXXXXX................XXXXXXXXXXXXXX..........XX
        XXX..............................................XXXXXXXXX...............XXXXXXXXXXXXXXX..........XX
        XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX.......XXXXXXXXX...............XXXXXXXXXXXXXXX..........XX
        XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX......XXXXXXXXX................XXXXXXXXXXXXXX..........XX
        XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX......XXXXXXXXX................XXXXXXXXXXXXXX..........XX
        XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX......XXXXXXXXX....XXXXXXXXX...XXXXXXXXXXXXXX..........XX
        XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX......XXXXXXXXX....XXXXXXXXXX..XXXXXXXXXXXXX...........XX
        XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX......XXXXXXXXX....XXXXXXXXXX..XXXXXXXXXXXX............XX
        XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX......XXXXXXXXX....XXXXXXXXXX....XXXXXXXXX.............XX
        XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX......XXXXXXXXX....XXXXXXXXXXX....XXXXXXX..............XX
        XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX......XXXXXXXXX....XXXXXXXXXX..........................XX
        XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX.......XXXXXXXXX....XXXXXXXXXX..........................XX
        XXX..............................................XXXXXXXXX....XXXXXXXXXX.......XXXXX..............XX
        XXX.......XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX....XXXXXXXXXX......XXXXXXXXX...........XX
        XXX......XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX....XXXXXXXXXX....XXXXXXXXXXX...........XX
        XXX.....XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX....XXXXXXXXXX...XXXXXXXXXXXXX..........XX
        XXX.....XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX....XXXXXXXXXX..XXXXXXXXXXXXXXX.........XX
        XXX.....XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX....XXXXXXXXXX..XXXXXXXXXXXXXXX.........XX
        XXX.....XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX....XXXXXXXXXX..XXXXXXXXXXXXXXX.........XX
        XXX.....XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX....XXXXXXXXXX..XXXXXXXXXXXXXXX.........XX
        XXX.....XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX....XXXXXXXXXX..XXXXXXXXXXXXXXX.........XX
        XXX.....XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX....XXXXXXXXXX..XXXXXXXXXXXXXXX.........XX
        XXX......XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX....XXXXXXXXXX..XXXXXXXXXXXXXXX.........XX
        XXX.......XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX.....XXXXXXXXXX...XXXXXXXXXXXXXX.........XX
        XXX...........................................................XXXXXXXXXX....XXXXXXXXXXX...........XX
        XXX...........................................................XXXXXXXXXX.....XXXXXXXXX............XX
        XXX...........................................................XXXXXXXXXX......XXXXXXX.............XX
        XXX...........................................................XXXXXXXXXX..........................XX
        XXX...........................................................XXXXXXXXXX..........................XX
        XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX...............XX
        XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX...............XX
        XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX...............XX
        XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX...............XX
        XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX...............XX
        XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX...............XX
        XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX...............XX
        XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX...............XX
        XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX...............XX
        XXX....XXXXXXXXXX....................XXXXXXXXXX..........................XXXXXXXXXX...............XX
        XXX....XXXXXXXXXX....................XXXXXXXXXX..........................XXXXXXXXXX...............XX
        XXX....XXXXXXXXXX....................XXXXXXXXXX..........................XXXXXXXXXX...............XX
        XXX....XXXXXXXXXX....................XXXXXXXXXX..........................XXXXXXXXXX...............XX
        XXX....XXXXXXXXXX....................XXXXXXXXXX..........................XXXXXXXXXX...............XX
        XXX....XXXXXXXXXX...XXXXXXX..........XXXXXXXXXX..........................XXXXXXXXXX...............XX
        XXX....XXXXXXXXXX...XXXXXXXX.........XXXXXXXXXX..........................XXXXXXXXXX..XXXXXX.......XX
        XXX....XXXXXXXXXX..XXXXXXXXXX........XXXXXXXXXX............XXXXXXXX......XXXXXXXXXX.XXXXXXXX......XX
        XXX....XXXXXXXXXX..XXXXXXXXXX........XXXXXXXXXX............XXXXXXXXX.....XXXXXXXXXXXXXXXXXXXX.....XX
        XXX....XXXXXXXXXX..XXXXXXXXXX........XXXXXXXXXX............XXXXXXXXX.....XXXXXXXXXXXXXXXXXXXXX....XX
        XXX....XXXXXXXXXX..XXXXXXXXXX........XXXXXXXXXX............XXXXXXXXX.....XXXXXXXXXXXXXXXXXXXXXX...XX
        XXX....XXXXXXXXXX..XXXXXXXXXX........XXXXXXXXXX............XXXXXXXXX.....XXXXXXXXXXXXXXXXXXXXXX...XX
        XXX....XXXXXXXXXX..XXXXXXXXXX........XXXXXXXXXX............XXXXXXXXX.....XXXXXXXXXXXXXXXXXXXXXX...XX
        XXX....XXXXXXXXXX..XXXXXXXXXX........XXXXXXXXXX............XXXXXXXXX.....XXXXXXXXXXXXXXXXXXXXXX...XX
        XXX....XXXXXXXXXX..XXXXXXXXXX.........XXXXXXXX.............XXXXXXXXX......XXXX....XXXXXXXXXXXXX...XX
        XXX....XXXXXXXXXX..XXXXXXXXXX..............................XXXXXXXXX...............XXXXXXXXXXX....XX
        XXX....XXXXXXXXXX..XXXXXXXXXX..............................XXXXXXXXX................XXXXXXXXX.....XX
        XXX....XXXXXXXXXX..XXXXXXXXXX..............................XXXXXXXXX................XXXXXXXX......XX
        XXX....XXXXXXXXXX..XXXXXXXXXX..............................XXXXXXXXX.................XXXXXX.......XX
        XXX....XXXXXXXXXX..XXXXXXXXXX..............................XXXXXXXXX..............................XX
        XXX....XXXXXXXXXX..XXXXXXXXXX..............................XXXXXXXXX..............................XX
        XXX....XXXXXXXXXX..XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX
        XXX....XXXXXXXXXX.XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX
        XXX....XXXXXXXXXX.XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX
        XXX....XXXXXXXXXX.XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX
        XXX....XXXXXXXXXX.XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX
        XXX....XXXXXXXXXX.XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX
        XXX....XXXXXXXXXX.XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX
        XXX....XXXXXXXXXX.XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX
        XXX....XXXXXXXXXX..XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX
        XXX....XXXXXXXXXX.................................................................................XX
        XXX....XXXXXXXXXX.................................................................................XX
        XXX....XXXXXXXXXX.................................................................................XX
        XXX....XXXXXXXXXX.................................................................................XX
        XXX....XXXXXXXXXX.................................................................................XX
        XXX....XXXXXXXXXX.................................................................................XX
        XXX....XXXXXXXXXX.................................................................................XX
        XXX....XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX...XX
        XXX....XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX..XX
        XXX....XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX..XX
        XXX....XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX..XX
        XXX....XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX..XX
        XXX....XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX..XX
        XXX....XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX..XX
        XXX....XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX..XX
        XXX....XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX..XX
        XXX....XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX..XX
        XXX....XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX...XX
        XXX.....XXXXXXXX..................................................................................XX
        XXX...............................................................................................XX
        XXX...............................................................................................XX
        XXX...............................................................................................XX
        XXX...............................................................................................XX
        XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX
        XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX
```