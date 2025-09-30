// Copyright (C) 2025 neocotic
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package problem

// Unwrapper is a function used by Builder.Wrap and Wrap to handle an already wrapped Problem in err's tree.
//
// An Unwrapper is effectively responsible for deciding what, if any, information from a wrapped Problem is to be used
// to construct the new Problem. Any such information will not take precedence over any explicitly defined Problem
// fields, however, it will take precedence over any information derived from a Definition or its Type.
type Unwrapper func(err error) Problem

// FullUnwrapper returns an Unwrapper that extracts all fields from a wrapped Problem in err's tree, if present. These
// fields will not take precedence over any explicitly defined Problem fields, however, it will take precedence over any
// fields derived from a Definition or its Type.
func FullUnwrapper() Unwrapper {
	return unwrapAllFields
}

// NoopUnwrapper returns an Unwrapper that does nothing.
func NoopUnwrapper() Unwrapper {
	return func(_ error) Problem {
		return Problem{}
	}
}

// PropagatedFieldUnwrapper returns an Unwrapper that extracts only fields that are expected to be propagated (e.g.
// captured stack trace, generated "UUID") from a wrapped Problem in err's tree, if present. Any such fields will not
// take precedence over any explicitly defined Problem fields, however, it will take precedence over any fields derived
// from a Definition or its Type.
func PropagatedFieldUnwrapper() Unwrapper {
	return unwrapPropagatedFields
}

// unwrapAllFields extracts all fields from a wrapped Problem in err's tree, if present. These fields will not take
// precedence over any explicitly defined Problem fields, however, it will take precedence over any fields derived from
// a Definition or its Type.
func unwrapAllFields(err error) Problem {
	if p, isProblem := As(err); isProblem && p != nil {
		return *p
	}
	return Problem{}
}

// unwrapPropagatedFields extracts only fields that are expected to be propagated (e.g. captured stack trace, generated
// "UUID") from a wrapped Problem in err's tree, if present. Any such fields will not take precedence over any
// explicitly defined Problem fields, however, it will take precedence over any fields derived from a Definition or its
// Type.
func unwrapPropagatedFields(err error) Problem {
	if p, isProblem := As(err); isProblem && p != nil {
		return Problem{
			Stack:   p.Stack,
			UUID:    p.UUID,
			logInfo: p.logInfo,
		}
	}
	return Problem{}
}
