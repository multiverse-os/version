package version

//import (
//	"errors"
//	"fmt"
//	"strconv"
//	"strings"
//)
//
//const (
//	numeric      string = "0123456789"
//	alphabetic          = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ-"
//	alphanumeric        = (numeric + alphabetic)
//)
//
//var DefaultSemanticVersion = SemanticVersion{
//	Major: 0,
//	Minor: 1,
//	Patch: 0,
//}
//
//type PRVersion string
//
//type SemanticVersion struct {
//	Major uint64
//	Minor uint64
//	Patch uint64
//	Pre   []PRVersion
//	Build []string //No Precedence
//}
//
//func (v SemanticVersion) String() string {
//	b := make([]byte, 0, 5)
//	b = strconv.AppendUint(b, v.Major, 10)
//	b = append(b, '.')
//	b = strconv.AppendUint(b, v.Minor, 10)
//	b = append(b, '.')
//	b = strconv.AppendUint(b, v.Patch, 10)
//
//	if len(v.Pre) > 0 {
//		b = append(b, '-')
//		b = append(b, v.Pre[0].String()...)
//
//		for _, pre := range v.Pre[1:] {
//			b = append(b, '.')
//			b = append(b, pre.String()...)
//		}
//	}
//
//	if len(v.Build) > 0 {
//		b = append(b, '+')
//		b = append(b, v.Build[0]...)
//
//		for _, build := range v.Build[1:] {
//			b = append(b, '.')
//			b = append(b, build...)
//		}
//	}
//
//	return string(b)
//}
//
//// FinalizeSemanticVersion discards prerelease and build number and only returns
//// major, minor and patch number.
//func (v SemanticVersion) FinalizeSemanticVersion() string {
//	b := make([]byte, 0, 5)
//	b = strconv.AppendUint(b, v.Major, 10)
//	b = append(b, '.')
//	b = strconv.AppendUint(b, v.Minor, 10)
//	b = append(b, '.')
//	b = strconv.AppendUint(b, v.Patch, 10)
//	return string(b)
//}
//
//func (self SemanticVersion) EQ(version SemanticVersion) bool { return self.Compare(version) == 0 }
//func (self SemanticVersion) NE(version SemanticVersion) bool { return self.Compare(version) != 0 }
//func (self SemanticVersion) GT(version SemanticVersion) bool { return self.Compare(version) == 1 }
//
//func (self SemanticVersion) GTE(SemanticVersion SemanticVersion) bool {
//	return self.Compare(version) >= 0
//}
//
//func (self SemanticVersion) GE(version SemanticVersion) bool {
//	return self.Compare(version) >= 0
//}
//
//func (self SemanticVersion) LT(version SemanticVersion) bool { return self.Compare(version) == -1 }
//
//// LTE checks if v is less than or equal to o.
//func (v SemanticVersion) LTE(o SemanticVersion) bool {
//	return (v.Compare(o) <= 0)
//}
//
//// LE checks if v is less than or equal to o.
//func (v SemanticVersion) LE(o SemanticVersion) bool {
//	return (v.Compare(o) <= 0)
//}
//
//// Compare compares SemanticVersions v to o:
//// -1 == v is less than o
//// 0 == v is equal to o
//// 1 == v is greater than o
//func (v SemanticVersion) Compare(o SemanticVersion) int {
//	if v.Major != o.Major {
//		if v.Major > o.Major {
//			return 1
//		}
//		return -1
//	}
//	if v.Minor != o.Minor {
//		if v.Minor > o.Minor {
//			return 1
//		}
//		return -1
//	}
//	if v.Patch != o.Patch {
//		if v.Patch > o.Patch {
//			return 1
//		}
//		return -1
//	}
//
//	// Quick comparison if a version has no prerelease versions
//	if len(v.Pre) == 0 && len(o.Pre) == 0 {
//		return 0
//	} else if len(v.Pre) == 0 && len(o.Pre) > 0 {
//		return 1
//	} else if len(v.Pre) > 0 && len(o.Pre) == 0 {
//		return -1
//	}
//
//	i := 0
//	for ; i < len(v.Pre) && i < len(o.Pre); i++ {
//		if comp := v.Pre[i].Compare(o.Pre[i]); comp == 0 {
//			continue
//		} else if comp == 1 {
//			return 1
//		} else {
//			return -1
//		}
//	}
//
//	// If all pr versions are the equal but one has further prversion, this one greater
//	if i == len(v.Pre) && i == len(o.Pre) {
//		return 0
//	} else if i == len(v.Pre) && i < len(o.Pre) {
//		return -1
//	} else {
//		return 1
//	}
//
//}
//
//// IncrementPatch increments the patch version
//func (v *SemanticVersion) IncrementPatch() error {
//	v.Patch++
//	return nil
//}
//
//// IncrementMinor increments the minor version
//func (v *SemanticVersion) IncrementMinor() error {
//	v.Minor++
//	v.Patch = 0
//	return nil
//}
//
//// IncrementMajor increments the major version
//func (v *SemanticVersion) IncrementMajor() error {
//	v.Major++
//	v.Minor = 0
//	v.Patch = 0
//	return nil
//}
//
//// Validate validates v and returns error in case
//func (v SemanticVersion) Validate() error {
//	// Major, Minor, Patch already validated using uint64
//
//	for _, pre := range v.Pre {
//		if !pre.IsNum { //Numeric prerelease versions already uint64
//			if len(pre.SemanticVersionStr) == 0 {
//				return fmt.Errorf("Prerelease can not be empty %q", pre.SemanticVersionStr)
//			}
//			if !containsOnly(pre.SemanticVersionStr, alphanum) {
//				return fmt.Errorf("Invalid character(s) found in prerelease %q", pre.SemanticVersionStr)
//			}
//		}
//	}
//
//	for _, build := range v.Build {
//		if len(build) == 0 {
//			return fmt.Errorf("Build meta data can not be empty %q", build)
//		}
//		if !containsOnly(build, alphanum) {
//			return fmt.Errorf("Invalid character(s) found in build meta data %q", build)
//		}
//	}
//
//	return nil
//}
//
//// New is an alias for Parse and returns a pointer, parses version string and returns a validated SemanticVersion or error
//func New(s string) (*SemanticVersion, error) {
//	v, err := Parse(s)
//	vp := &v
//	return vp, err
//}
//
//// Make is an alias for Parse, parses version string and returns a validated SemanticVersion or error
//func Make(s string) (SemanticVersion, error) {
//	return Parse(s)
//}
//
//// ParseTolerant allows for certain version specifications that do not strictly adhere to semver
//// specs to be parsed by this library. It does so by normalizing versions before passing them to
//// Parse(). It currently trims spaces, removes a "v" prefix, adds a 0 patch number to versions
//// with only major and minor components specified, and removes leading 0s.
//func ParseTolerant(s string) (SemanticVersion, error) {
//	s = strings.TrimSpace(s)
//	s = strings.TrimPrefix(s, "v")
//
//	// Split into major.minor.(patch+pr+meta)
//	parts := strings.SplitN(s, ".", 3)
//	// Remove leading zeros.
//	for i, p := range parts {
//		if len(p) > 1 {
//			p = strings.TrimLeft(p, "0")
//			if len(p) == 0 || !strings.ContainsAny(p[0:1], "0123456789") {
//				p = "0" + p
//			}
//			parts[i] = p
//		}
//	}
//	// Fill up shortened versions.
//	if len(parts) < 3 {
//		if strings.ContainsAny(parts[len(parts)-1], "+-") {
//			return SemanticVersion{}, errors.New("Short version cannot contain PreRelease/Build meta data")
//		}
//		for len(parts) < 3 {
//			parts = append(parts, "0")
//		}
//	}
//	s = strings.Join(parts, ".")
//
//	return Parse(s)
//}
//
//// Parse parses version string and returns a validated SemanticVersion or error
//func Parse(s string) (SemanticVersion, error) {
//	if len(s) == 0 {
//		return SemanticVersion{}, errors.New("SemanticVersion string empty")
//	}
//
//	// Split into major.minor.(patch+pr+meta)
//	parts := strings.SplitN(s, ".", 3)
//	if len(parts) != 3 {
//		return SemanticVersion{}, errors.New("No Major.Minor.Patch elements found")
//	}
//
//	// Major
//	if !containsOnly(parts[0], numbers) {
//		return SemanticVersion{}, fmt.Errorf("Invalid character(s) found in major number %q", parts[0])
//	}
//	if hasLeadingZeroes(parts[0]) {
//		return SemanticVersion{}, fmt.Errorf("Major number must not contain leading zeroes %q", parts[0])
//	}
//	major, err := strconv.ParseUint(parts[0], 10, 64)
//	if err != nil {
//		return SemanticVersion{}, err
//	}
//
//	// Minor
//	if !containsOnly(parts[1], numbers) {
//		return SemanticVersion{}, fmt.Errorf("Invalid character(s) found in minor number %q", parts[1])
//	}
//	if hasLeadingZeroes(parts[1]) {
//		return SemanticVersion{}, fmt.Errorf("Minor number must not contain leading zeroes %q", parts[1])
//	}
//	minor, err := strconv.ParseUint(parts[1], 10, 64)
//	if err != nil {
//		return SemanticVersion{}, err
//	}
//
//	v := SemanticVersion{}
//	v.Major = major
//	v.Minor = minor
//
//	var build, prerelease []string
//	patchStr := parts[2]
//
//	if buildIndex := strings.IndexRune(patchStr, '+'); buildIndex != -1 {
//		build = strings.Split(patchStr[buildIndex+1:], ".")
//		patchStr = patchStr[:buildIndex]
//	}
//
//	if preIndex := strings.IndexRune(patchStr, '-'); preIndex != -1 {
//		prerelease = strings.Split(patchStr[preIndex+1:], ".")
//		patchStr = patchStr[:preIndex]
//	}
//
//	if !containsOnly(patchStr, numbers) {
//		return SemanticVersion{}, fmt.Errorf("Invalid character(s) found in patch number %q", patchStr)
//	}
//	if hasLeadingZeroes(patchStr) {
//		return SemanticVersion{}, fmt.Errorf("Patch number must not contain leading zeroes %q", patchStr)
//	}
//	patch, err := strconv.ParseUint(patchStr, 10, 64)
//	if err != nil {
//		return SemanticVersion{}, err
//	}
//
//	v.Patch = patch
//
//	// Prerelease
//	for _, prstr := range prerelease {
//		parsedPR, err := NewPRSemanticVersion(prstr)
//		if err != nil {
//			return SemanticVersion{}, err
//		}
//		v.Pre = append(v.Pre, parsedPR)
//	}
//
//	// Build meta data
//	for _, str := range build {
//		if len(str) == 0 {
//			return SemanticVersion{}, errors.New("Build meta data is empty")
//		}
//		if !containsOnly(str, alphanum) {
//			return SemanticVersion{}, fmt.Errorf("Invalid character(s) found in build meta data %q", str)
//		}
//		v.Build = append(v.Build, str)
//	}
//
//	return v, nil
//}
//
//// MustParse is like Parse but panics if the version cannot be parsed.
//func MustParse(s string) SemanticVersion {
//	v, err := Parse(s)
//	if err != nil {
//		panic(`semver: Parse(` + s + `): ` + err.Error())
//	}
//	return v
//}
//
//// PRSemanticVersion represents a PreRelease SemanticVersion
//type PRSemanticVersion struct {
//	SemanticVersionStr string
//	SemanticVersionNum uint64
//	IsNum              bool
//}
//
//// NewPRSemanticVersion creates a new valid prerelease version
//func NewPRSemanticVersion(s string) (PRSemanticVersion, error) {
//	if len(s) == 0 {
//		return PRSemanticVersion{}, errors.New("Prerelease is empty")
//	}
//	v := PRSemanticVersion{}
//	if containsOnly(s, numbers) {
//		if hasLeadingZeroes(s) {
//			return PRSemanticVersion{}, fmt.Errorf("Numeric PreRelease version must not contain leading zeroes %q", s)
//		}
//		num, err := strconv.ParseUint(s, 10, 64)
//
//		// Might never be hit, but just in case
//		if err != nil {
//			return PRSemanticVersion{}, err
//		}
//		v.SemanticVersionNum = num
//		v.IsNum = true
//	} else if containsOnly(s, alphanum) {
//		v.SemanticVersionStr = s
//		v.IsNum = false
//	} else {
//		return PRSemanticVersion{}, fmt.Errorf("Invalid character(s) found in prerelease %q", s)
//	}
//	return v, nil
//}
//
//// IsNumeric checks if prerelease-version is numeric
//func (v PRSemanticVersion) IsNumeric() bool {
//	return v.IsNum
//}
//
//// Compare compares two PreRelease SemanticVersions v and o:
//// -1 == v is less than o
//// 0 == v is equal to o
//// 1 == v is greater than o
//func (v PRSemanticVersion) Compare(o PRSemanticVersion) int {
//	if v.IsNum && !o.IsNum {
//		return -1
//	} else if !v.IsNum && o.IsNum {
//		return 1
//	} else if v.IsNum && o.IsNum {
//		if v.SemanticVersionNum == o.SemanticVersionNum {
//			return 0
//		} else if v.SemanticVersionNum > o.SemanticVersionNum {
//			return 1
//		} else {
//			return -1
//		}
//	} else { // both are Alphas
//		if v.SemanticVersionStr == o.SemanticVersionStr {
//			return 0
//		} else if v.SemanticVersionStr > o.SemanticVersionStr {
//			return 1
//		} else {
//			return -1
//		}
//	}
//}
//
//// PreRelease version to string
//func (v PRSemanticVersion) String() string {
//	if v.IsNum {
//		return strconv.FormatUint(v.SemanticVersionNum, 10)
//	}
//	return v.SemanticVersionStr
//}
//
//func containsOnly(s string, set string) bool {
//	return strings.IndexFunc(s, func(r rune) bool {
//		return !strings.ContainsRune(set, r)
//	}) == -1
//}
//
//func hasLeadingZeroes(s string) bool {
//	return len(s) > 1 && s[0] == '0'
//}
//
//// NewBuildSemanticVersion creates a new valid build version
//func NewBuildSemanticVersion(s string) (string, error) {
//	if len(s) == 0 {
//		return "", errors.New("Buildversion is empty")
//	}
//	if !containsOnly(s, alphanum) {
//		return "", fmt.Errorf("Invalid character(s) found in build meta data %q", s)
//	}
//	return s, nil
//}
//
//// FinalizeSemanticVersion returns the major, minor and patch number only and discards
//// prerelease and build number.
//func FinalizeSemanticVersion(s string) (string, error) {
//	v, err := Parse(s)
//	if err != nil {
//		return "", err
//	}
//	v.Pre = nil
//	v.Build = nil
//
//	finalVer := v.String()
//	return finalVer, nil
//}
