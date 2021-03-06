package heap

import "jvmgo/src/classfile"

type MethodRef struct {
	MemberRef
	method *Method
}

func newMethodRef(cp *ConstantPool, refInfo *classfile.ConstantMethodrefInfo) *MethodRef {
	ref := &MethodRef{}
	ref.cp = cp
	ref.copyMemberRefInfo(&refInfo.ConstantMemberrefInfo)

	return ref
}


func (self *MethodRef) ResolvedMethod() *Method {
	if self.method == nil {
		self.resolveMethodRef()
	}
	return self.method
}

func (self *MethodRef) resolveMethodRef() *Method {
	d := self.cp.class
	c := self.ResolvedClass()

	if c.IsInterface() {
		panic("java.lang.IncompatibleCLassChangeError")
	}

	method := lookupMethod(c, self.name, self.descriptor)
	if method == nil {
		panic("java.lang.NoSuchMethodError")
	}

	if !method.IsAccessibleTo(d) {
		panic("java.lang.IllegalAccessError")
	}

	return method
}

func lookupMethod(class *Class, name string, descriptor string) *Method {
	method := LookupMethodImClass(class, name, descriptor)
	if method == nil {
		method = lookupMethodInterfaces(class.interfaces, name, descriptor)
	}
	
	return method
}
func lookupMethodInterfaces(ifaces []*Class, name string, descriptor string) *Method {
	for _,iface := range ifaces {
		for _, method := range iface.methods {
			if method.name == name && method.descriptor == descriptor {
				return method
			}
		}
		method := lookupMethodInterfaces(iface.interfaces, name, descriptor)
		if method != nil {
			return method
		}
	}

	return nil
}
func LookupMethodImClass(class *Class, name string, descriptor string) *Method {
	for c := class; c != nil; c = c.superClass {
		for _, method := range c.methods {
			if method.name == name && method.descriptor == descriptor {
				return method
			}
		}
	}
	
	return nil
}