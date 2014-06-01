FileStat = ->
    self = this
    self.read = true
    self.write = true

    return

Dir = ->
    self = this
    self.stat = new FileStat()
    self.entries = {}

    self.add = (name, obj) ->
        if !self.stat.read || !self.stat.write
            return 'permission denied'
        if name of self.entries
            return 'already exists'
        self.entries[name] = obj
        return ''

    self.rm = (name) ->
        if !self.stat.read || !self.stat.write
            return 'permission denied'
        if !(name of self.entries)
            return 'not exists'
        delete self.entries[name]
        return ''

    self.get = (name) ->
        if !self.stat.read
            return 'permission denied'
        if !(name of self.entries)
            return 'not exists'
        return self.entries[name]

    self.have = (name) -> name of self.entries

    self.list = ->
        ret = []
        for k of self.entries
            ret.push(k)
        return ret

    self.isDir = true
    self.isFile = false

    return

File = ->
    self = this
    self.stat = new FileStat()

    self.isDir = false
    self.isFile = true
    
    return

isString = (s) -> s.constructor == String

FileSys = ->
    self = this
    self.root = new Dir()

    pathBreak = (p) -> p.substr(1).split('/')

    # open is for read
    self.open = (path) ->
        parts = pathBreak(path)
        ret = self.root
        for p in parts
            if !ret.isDir
                return 'not found'
            ret = ret.get(p)
            if isString(ret)
                return ret

        return ret

    # create is for write
    self.create = (path) ->
        parts = pathBreak(path)
        npart = parts.length
        dirs = parts.slice(0, npart-1)
        file = parts[npart-1]
        ret = self.root
        for p in dirs
            ret = ret.get(p)
            if isString(ret)
                return ret
            if !ret.isDir
                return 'dir not found'
        
        if ret.have(file)
            ret.rm(file)

        f = new File()
        e = ret.add(file, f)
        if e != ''
            return e

        return f

    self.lsdir = (path) ->
        parts = pathBreak(path)
        ret = self.root
        for p in parts
            ret = ret.get(p)
            if isString(rret)
                return ret
            if !ret.isDir
                return 'not dir'
        return ret.list()

    self.mkdir = (path) ->
        return

    self.rmdir = (path) ->
        return

    self.mv = (from, to) ->
        return

    self.rm = (path) ->
        return

    return

