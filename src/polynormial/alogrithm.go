package polynormail

import (
	"fmt"
	"math/big"
	"math/rand"
)

const PrimeSize = 10000

type Poly struct {
	p big.Int
}

var polyRand *rand.Rand
var Prime []*Poly

func init() {
	polyRand = rand.New(rand.NewSource(1))
	Prime = FindPrime(18)
}

func NewXn(n int) *Poly {
	if n < 0 {
		fmt.Println("Xn: must n >= 0", n)
	}
	r := big.NewInt(1)
	for i := 0; i < n; i++ {
		r.Mul(r, big.NewInt(2))
	}
	return &Poly{p: *r}
}

func NewRand(n int) *Poly {
	m := NewXn(n).p
	v := big.NewInt(0)
	v.Rand(polyRand, &m)
	return &Poly{p: *v}
}

func (x *Poly) Println(s string) {
	fmt.Println(s, (&x.p).Text(16))
}

func (x *Poly) Order() int {
	y := x.p
	i := 0
	r := big.NewInt(1)
	for {
		if r.Cmp(&y) <= 0 {
			i += 1
			r.Mul(r, big.NewInt(2))
		} else {
			break
		}

	}
	return i
}

func (x *Poly) PrintPoly() {
	o := x.Order()
	y := &(x.p)
	for i, j := o, 0; i > 0; i-- {
		if y.Bit(i) != 0 {
			fmt.Printf("X[%03d]", i)
			j++
			if j == 8 {
				j = 0
				fmt.Printf("\n")
			}
		}
	}
	if y.Bit(0) != 0 && o != 0 {
		fmt.Printf("1")
	}
	fmt.Println()
}

func NewPoly(y *Poly) *Poly {
	x := big.NewInt(0)
	z := y.p
	x.SetBytes((&z).Bytes())
	return &Poly{p: *x}
}

func (z *Poly) Add(x, y *Poly) *Poly {
	r := (&z.p).Xor(&x.p, &y.p)
	z.p = *r
	return z
}

func (z *Poly) Mul(x, y *Poly) *Poly {
	s := big.NewInt(0)
	o := y.Order()
	vx := NewPoly(x)
	ix, iy := vx.p, y.p
	px, py := &ix, &iy
	for i := 0; i <= o; i++ {
		if py.Bit(i) != 0 {
			s.Xor(s, px)
		}
		px.Lsh(px, 1)
	}
	z.p = *s
	return z
}

func (a *Poly) DivRem(b *Poly) *Poly {
	oa := a.Order()
	ob := b.Order()
	r := big.NewInt(0)
	d := big.NewInt(0)
	if oa < ob {
		return &Poly{p: *r}
	}
	or := uint(oa - ob)
	d.Lsh(&b.p, or)
	pa := &a.p
	for i := 0; i < int(or+1); i++ {
		r.Lsh(r, 1)
		if pa.Bit(oa-i-1) != 0 {
			pa.Xor(pa, d)
			r.Xor(r, big.NewInt(1))
		}
		d.Rsh(d, 1)
	}
	return &Poly{p: *r}
}

func FindPrime(n int) []*Poly {
	max := &(NewXn(n/2 + 1).p)
	r := make([]*Poly, PrimeSize)
	cnt := 1
	r[0] = &Poly{p: *big.NewInt(2)}
	i := int64(3)
	for i < (int64(1) << uint(n)) {
		find := 0
		for _, d := range r[:cnt] {
			s := &Poly{p: *big.NewInt(i)}
			// s.Println("as")
			s.DivRem(d)
			// s.Println("bs")
			// d.Println("d")
			if (&s.p).Sign() == 0 {
				find = 1
				break
			}
			if max.Cmp(&d.p) < 0 {
				break
			}
		}
		if find == 0 {
			if cnt < PrimeSize {
				r[cnt] = &Poly{p: *big.NewInt(i)}
				if cnt%1024 == 0 {
					r[cnt].Println(fmt.Sprintf("%2d", cnt/1024))
				}
				cnt++
			} else {
				break
			}
		}
		i++
	}
	return r[:cnt]
}

func (a *Poly) Factorize() []*Poly {
	d := NewPoly(a)
	r := make([]*Poly, PrimeSize)
	cnt := 0
	for _, p := range Prime {
		for {
			s := NewPoly(d)
			m := s.DivRem(p)
			if (&s.p).Sign() == 0 {
				r[cnt] = p
				cnt++
				d = m
			} else {
				break
			}
		}
		if d.Order() < p.Order() {
			break
		}
	}
	if (&d.p).Sign() != 0 {
		r[cnt] = d
		cnt++
	}
	return r[:cnt]
}

// func (a *Poly) Rdiv(b *Poly) (*Poly, *Poly) {
//     dl
// }

/*

#include"strPoly.h"
#include<memory.h>
#include<stdio.h>
#include<assert.h>




strPoly::strPoly()
{
    length=0;
    debug=0;
    poly = new UCHAR[POLYSIZE];
    memset( poly, 0, POLYSIZE );
    div=NULL;
    rem=NULL;
    index=NULL;
    power=NULL;
    mulTab=NULL;
    pRoot=NULL;
}

strPoly::~strPoly( )
{
    delete[] poly;
    if( div )
        delete[] div;
    if( rem )
        delete[] rem;
    if( index )
        delete[] index;
    if( power )
        delete[] power;
    if( mulTab )
        delete[] mulTab;
}

int strPoly::genTable()
{
    if( length>9 )
        return -1;
    div = new UCHAR[256];
    rem = new UCHAR[256];
    index = new UCHAR[256];
    power = new UCHAR[256];

    strPoly pa,pb;
    int i,j,k;
    rem[0]=div[0]=0;
    for( i=0;i<256;i++ )
    {
        pa.int2Poly( i<<(length-1) );
#if 0
        printf("a:");
        pa.printPoly();
        printf("\n");
        printf("p:");
        printPoly();
        printf("\n");
#endif
        pb.pdiv(pa,*this);
#if 0
        printf("r:");
        pa.printPoly();
        printf("\n");
        printf("d:");
        pb.printPoly();
        printf("\n");
#endif
        j=i<<(length-1);
        j=(j&0xff)|(j>>8);
        j&=0xff;
        if( (k=pa.poly2Int( rem+j ))>1 )
        {
            printf("rem error [len=%d]:",k);
            pa.printPoly();
            printf("\n");
            return -2;
        }
        if( pb.poly2Int( div+j )>1 )
            return -3;

    }
    order = (1<<(length-1))-1;
    int a=1;
    int gfp=0;
    if( poly2Int((UCHAR *)&gfp)>2 )
        return -4;
#if 0
    printf("p=%x\n",gfp);
    printPoly();
    printf("\n");
#endif
    for( i=0;i<order;i++ )
    {
        power[i]=a;
        index[a]=i;
        a<<=1;
        if( (a>>(length-1))&1 )
            a^=gfp;
    }
    index[0]=order;
    power[order]=0;
#if 0
    for( i=0;i<order+1;i++ )
    {
        printf("index[%d]=%d\n",i,index[i]);
    }
#endif
#if 0
    for( i=0;i<256;i++ )
    {
        printf("div[%d]=%02x rem[%d]=%02x\n",i,div[i]&0xff,i,rem[i]&0xff);
    }
#endif

    return 0;
}

void strPoly::printPoly( )
{

    int i;
    int j=0;
    for( i=length-1;i>0;i-- )
    {
        if(poly[i]!=0)
        {
            printf("X[%03d]",i);
            j++;
            if( j==8 )
            {
                j=0;
                printf("\n");
            }
        }
    }
    if( poly[0]!=0 && length!=0 )
        printf("%x",poly[0]);
}

void strPoly::copyPoly( strPoly& a )
{
    this->length=a.length;
    int i;
    memset( poly, 0, POLYSIZE );
    for( i=0;i<this->length;i++ )
        this->poly[i]=a.poly[i];
}

void strPoly::xk( int k )
{
    this->length=k+1;
    memset( poly, 0, POLYSIZE );
    this->poly[k]=1;
}

void strPoly::pdiv( strPoly& a, strPoly& b )
{

    int i,j;
    memset( poly, 0, POLYSIZE );
    this->length=a.length+1-b.length;
    assert( b.poly[b.length-1]==1 );
    for( i=a.length-1;i>=b.length-1;i-- )
    {
        this->poly[i-b.length+1]=a.poly[i];
        if( a.poly[i]==1 )
            for( j=0;j<b.length-1;j++ )
                a.poly[i-j-1]^=b.poly[b.length-2-j];
    }
    a.length=b.length-1;
    for( i=a.length-1;i>=0;i-- )
    {
        if( a.poly[i]!=0 )
        {
            a.length=i+1;
            break;
        }
    }
    if( i==-1 )
        a.length=0;
}

void strPoly::padd( strPoly& a, strPoly& b )
{
    strPoly d;
    int i,j;
    if( a.length>b.length )
    {
        d.copyPoly(a);
        for( i=0;i<b.length;i++ )
            d.poly[i]^=b.poly[i];
    }
    else
    {
        d.copyPoly(b);
        for( i=0;i<a.length;i++ )
            d.poly[i]^=a.poly[i];
    }
    for( i=d.length-1;i>=0;i-- )
    {
        if( d.poly[i]!=0 )
        {
            d.length=i+1;
            break;
        }
    }
    if( i==-1 )
        d.length=0;
    this->copyPoly(d);
}


void strPoly::pmul( strPoly& a, strPoly& b )
{

    int i,j;
    strPoly d;
    d.length=a.length+b.length-1;
    if( debug==2 )
    {
        printf("\nMul: a");
        a.printPoly();
        printf("\nMul: b");
        b.printPoly();

    }
    for( i=0;i<d.length;i++)
        d.poly[i]=0;
    for( i=0;i<b.length;i++ )
    {
        for( j=0;j<a.length;j++ )
            d.poly[i+j]^=(a.poly[j]&b.poly[i]);
    }
    if( debug==2 )
    {
        printf("\nMul: c");
        d.printPoly();
        printf("\n");

    }
    copyPoly(d);
}

void strPoly::int2Poly( int i )
{
    int k;
    k=i;
    length=0;
    for(;k!=0;k>>=1)
    {
        poly[length]=k&1;
        length++;
#if 0
        printf("%d: ",k);
        printPoly();
        printf("\n");
#endif
    }
}

int strPoly::findp( strPoly dpp[], int p )
{
    strPoly a,b,c,d;
    int i,j,l,k;

    a.xk(p);
    a.poly[0]=1;

    i=3;
    l=0;

    while(i<0x80000000)
    {
        b.int2Poly(i);

        for(k=0;k<l;k++)
        {
            d.copyPoly(b);
            c.pdiv(d,dpp[k]);
            if( c.length==0 )
                break;
        }

        while( k==l )
        {
            d.copyPoly(a);
            c.pdiv(d,b);
            if( d.length!=0 )
                break;
            dpp[l].copyPoly(b);
            l++;
            a.copyPoly(c);
        }
        if( a.length<b.length )
            break;
        i++;
    }
    return l;
}

int strPoly::rdiv( strPoly& a, strPoly& b, strPoly& m, strPoly& n)
{

    if( debug==1 )
    {
        printf("\n a:\n");
        a.printPoly();
        printf("\n b:\n");
        b.printPoly();
    }
    strPoly *dl = new strPoly[1024];
    strPoly g1,g2,g3;
    strPoly *pg1,*pg2,*pg3;
    strPoly *mk0,*mk1,*mk2;
    int i,j,k;

    g3.copyPoly(b);
    g2.copyPoly(a);
    pg1=&g1;
    pg2=&g2;
    pg3=&g3;
    i=0;
    strPoly *t;

    while( pg3->length != 1 )
    {
        t=pg1;
        pg1=pg2;
        pg2=pg3;
        pg3=t;
        if( debug==1 )
        {
            printf("\ndiv1:");
            pg1->printPoly();
            printf("\ndiv1:");
            pg2->printPoly();
        }
        dl[i].pdiv(*pg1,*pg2);
        if( debug==1 )
        {
            printf("\nd[%d]:",i);
            dl[i].printPoly();
            printf("\nrem:");
            pg1->printPoly();
            printf("\n");
        }
        pg3->copyPoly(*pg1);
        i++;
        if( pg3->length==0 )
        {
            m.copyPoly(*pg2);
            return -1;
        }
    }

    mk2=&g3;
    mk1=&g2;
    mk0=&g1;

    mk2->length=1;
    mk2->poly[0]=1;

    mk1->copyPoly( dl[i-1] );
    n.copyPoly( *mk2 );
    m.copyPoly( *mk1 );

    for( j=i-2;j>=0;j-- )
    {
        if( debug==1 )
        {
            printf("\nmk2:");
            mk2->printPoly();
            printf("\nmk1:");
            mk1->printPoly();
            printf("\nl[%d]:",j);
            dl[j].printPoly();
            printf("\n");
        }
        mk0->pmul( *mk1,dl[j]);
        if( debug==1 )
        {
            printf("\nbefore add mk0:");
            mk0->printPoly();
            printf("\n");
        }
        mk0->padd( *mk0, *mk2 );
        if( debug==1 )
        {
            printf("\nmk0:");
            mk0->printPoly();
            printf("\n");
        }
        n.copyPoly( *mk1 );
        m.copyPoly( *mk0 );
        t=mk2;
        mk2=mk1;
        mk1=mk0;
        mk0=t;
    }

    if( debug==1 )
    {
        printf("\n m:\n");
        m.printPoly( );
        printf("\n n:\n");
        n.printPoly( );
    }

    delete[] dl;
    return 0;

}

int strPoly::poly2Int( UCHAR *p )
{
    int i,j,k;
    if( length==0 )
        p[0]=0;
    for( j=0;j<length;j+=8 )
    {
        p[j/8]=0;
        for( k=0;(k<8)&&((k+j)<length);k++ )
            if(poly[j+k]==1)
                p[j/8]|=(1<<k);
    }
    return (length+7)/8;
}

void strPoly::analyze( strPoly b[], int num, strPoly dp[] )
{
    int i;
    strPoly c;
    for( i=0;i<num;i++ )
    {
        b[i].copyPoly(*this);
        c.pdiv(b[i],dp[i]);
    }
}


void strPoly::synthesize( strPoly a[], int num, int p, strPoly dp[], strPoly mp[] )
{
    int i;
    strPoly c;
    for( i=0;i<num;i++ )
    {
        a[i].pmul(a[i],mp[i]);
        c.pdiv(a[i],dp[i]);
#if DEBUGSYN
        printf("a[%d] :",i);
        a[i].printPoly();
        printf("\n");
#endif
    }
    length=0;
    for( i=0;i<num;i++ )
    {
        c.xk(p);
        c.poly[0]=1;
        c.pmul(a[i],c);
        a[i].pdiv(c,dp[i]);
        padd(a[i],*this);
    }

}

UCHAR strPoly::remainder( UCHAR *p, UCHAR *d, int len )
{
    UCHAR r=0;
    UCHAR s=0;
    int i;
    UCHAR mask=(1<<(length-1))-1;
//  printf("mask:%02x\n",(int)mask&0xff);
    s=p[len-1]&(~mask);
    r=p[len-1]&(mask);
    for( i=len-1;i>=0;i-- )
    {
        d[i]=div[s];
        r^=rem[s];
#if 0
        printf(" s=%02x r=%02x div[s]=%02x rem[s]=%02x d[%d]=%02x p[i]=%02x p[i-1]=%02x\n",
            s&0xff,r&0xff,div[s]&0xff,rem[s]&0xff,i,d[i]&0xff,p[i]&0xff,p[i-1]&0xff);
#endif
        if( i!=0 )
        {
            s=r|(p[i-1]&(~mask));
            r=p[i-1]&(mask);
        }
    }

    return r;
}

int strPoly::genMulTab( strPoly& p )
{
    int i,j,k;
    mulTab = new UCHAR[256];
    pRoot = &p;
    mulTab[0]=0;
    int myIndex=0;
    if( poly2Int( (UCHAR *)&myIndex )> 1 )
        return -1;
#if 0
    printf("A-myIndex=%02x %02x\n",myIndex,p.index[myIndex]);
    printf(" index:\n");
    for( i=0;i<256;i++ )
    {
        printf("%02x ",p.index[i]&0xff);
        if( (i%16)==15 )
            printf("\n");
    }

    p.printPoly();
    printf("\n");
#endif
    if( myIndex==0 )
    {
        for( i=1;i<p.order;i++ )
            mulTab[i]=0;
        return 0;
    }
    myIndex = p.index[myIndex];
    for( i=1;i<=p.order;i++ )
    {
        int k=p.index[i];
        k+=myIndex;
        k%=p.order;
        mulTab[i]=p.power[k];
    }
    if( p.order==1 )
        mulTab[p.order]=1;
    return p.order;
}

UCHAR strPoly::gmul( UCHAR i, UCHAR j )
{

    if( i==0 || j==0 )
        return 0;
    int res=index[i]+index[j];
    res%=order;
    return power[res];

}

UCHAR strPoly::ginv( UCHAR i )
{

    if( i==0 )
        return 0;
    int res=order-index[i];
    res%=order;
    return power[res];

}


#include"strPoly.h"
#include<memory.h>
#include<stdio.h>
#include<assert.h>
#define EXT 25
extern "C" int GP256_inv( int *a, int *b )
{
    strPoly ap,ta,x8,m,n,g256;
    UCHAR rb[32];
    int i;
    for( i=0;i<8;i++ )
    {
        ta.int2Poly(a[i]);
        x8.xk(i*32);
        ta.pmul(ta,x8);
        ap.padd(ap,ta);
    }
    g256.xk(256);
    g256.poly[0]=1;
    int result=ta.rdiv(ap,g256,m,n);
    memset(rb,0,32);
    m.poly2Int(rb);
    memcpy(b,rb,32);
    ta.pmul(ap,m);
    ta.printPoly();
    return result;
}

extern "C" int GP256_mul( int *a, int *b, int *c )
{
    strPoly ap,bp,ta,tb,x8,m,n,g256;
    UCHAR rb[32];
    int i;
    for( i=0;i<8;i++ )
    {
        ta.int2Poly(a[i]);
        x8.xk(i*32);
        ta.pmul(ta,x8);
        ap.padd(ap,ta);
        tb.int2Poly(b[i]);
        x8.xk(i*32);
        tb.pmul(tb,x8);
        bp.padd(bp,ta);
    }
    ap.pmul(ap,bp);
    g256.xk(256);
    g256.poly[0]=1;
    ap.pdiv(ap,g256);
    memset(rb,0,32);
    ap.poly2Int(rb);
    memcpy(c,rb,32);
    ap.printPoly();
    return ap.length;
}

static void printMetrix( strPoly *m, int row, int col, int sr, int lr, int sc, int lc )
{

    int i,j;
    for( i=sc;i<sc+lc;i++ )
    {
        for( j=sr;j<sr+lr;j++ )
            printf("%3d ",m[i*row+j].length);
        printf("\n");
    }
    printf("\n");
}

static int polyCheck( strPoly &a )
{
    int r=0;
    for( int i=0;i<a.length;i++ )
        r^=a.poly[i];
    return r;
}

static int g256_inv( strPoly &a, strPoly &b )
{
    strPoly ap,ta,m,n,g256,v;

    g256.xk(256);
    g256.poly[0]=1;
    int result=ta.rdiv(a,g256,m,n);
    if( result<0 )   printf("rdiv error\n");
    ta.pmul(a,n);
    v.pdiv(ta,g256);
    ta.printPoly();
    b.copyPoly(n);
    return result;

}

extern "C" int GP256_matrix_guass( int c[], int row, int col, int bl, int pos, unsigned char *inv, unsigned char *ninv )
{
    strPoly *metrix = new strPoly[(row+col)*col];
    int len = row+col;
    int i,j,k;
    int count=0;
    strPoly g256;
    g256.xk(256);
    g256.poly[0]=1;
    for( i=0;i<col;i++ )
    {
        for( j=0;j<row;j++ )
            if( c[i*row+j]!=bl )
                metrix[i*len+j].xk(c[i*row+j]);
        metrix[i*len+row+i].xk(0);
        count++;
        if( (count%1000) == 0 )
            printf("%d\n",count);
    }

    for( i=0;i<pos;i++ )
    {
        for( j=i+1;j<col;j++ )
        {
            strPoly t,a,v;
            t.copyPoly(metrix[j*len+i]);
            if( t.length==0 ) continue;
            for( k=0;k<col+EXT;k++ )
            {
                a.debug=0;
                if( metrix[i*len+k].length==0 ) continue;
                a.pmul(metrix[i*len+k],t);
                v.pdiv(a,g256);
                metrix[j*len+k].padd(metrix[j*len+k],a);
                count++;
                if( (count%1000) == 0 )
                    printf("%d %d %d %d\n",count,i,j,k);
            }
        }
    }
printMetrix( metrix, len, col, 0, col, 0, col );
    for( i=0;i<col;i++ )
    {
        for( j=pos;j<col;j++ )
        {
            metrix[i*len+j].poly2Int(ninv);
            ninv+=32;
        }
    }

    for( i=pos;i<col;i++ )
    {
        strPoly t,ii,invii;
        for( j=i;j<col+EXT;j++ )
        {
            if( polyCheck(metrix[i*len+j])!=0 )
                break;
        }
        if( j==col+EXT )
        {
            printf("metrix can not inv\n");
            return -1;
        }
        if( j!=i )
        {
            printf("swap %d %d\n",i,j);
            for( int si=0;si<col;si++ )
            {
                int ti=c[si*row+i];
                c[si*row+i]=c[si*row+j];
                c[si*row+j]=ti;
                t.copyPoly(metrix[si*len+i]);
                metrix[si*len+i].copyPoly(metrix[si*len+j]);
                metrix[si*len+j].copyPoly(t);
            }
        }
        ii.copyPoly(metrix[i*len+i]);
        g256_inv(ii,invii);
        for( j=0;j<len;j++ )
        {
            metrix[i*len+j].pmul(metrix[i*len+j],invii);
            t.pdiv(metrix[i*len+j],g256);
                count++;
                if( (count%1000) == 0 )
                    printf("%d\n",count);
        }
        printf("after inv\n");
        printMetrix( metrix, len, col, 0, col, 0, col );
        for( j=i+1;j<col;j++ )
        {
            strPoly t,a,v;
            t.copyPoly(metrix[j*len+i]);
            if( t.length==0 ) continue;
            for( k=0;k<len;k++ )
            {
                a.debug=0;
                a.pmul(metrix[i*len+k],t);
                v.pdiv(a,g256);
                metrix[j*len+k].padd(metrix[j*len+k],a);
                count++;
                if( (count%1000) == 0 )
                    printf("%d\n",count);
            }
        }
        printMetrix( metrix, len, col, 0, col, 0, col );
    }
    for( i=col-1;i>=pos;i-- )
    {
        for( j=pos;j<i;j++ )
        {
            strPoly t,a,v;
            t.copyPoly(metrix[j*len+i]);
            for( k=0;k<len;k++ )
            {
                a.debug=0;
                a.pmul(metrix[i*len+k],t);
                v.pdiv(a,g256);
                metrix[j*len+k].padd(metrix[j*len+k],a);
                count++;
                if( (count%1000) == 0 )
                    printf("%d\n",count);
            }
        }
        printMetrix( metrix, len, col, 0, col, 0, col );
    }
    printf("inv metrix\n");
        printMetrix( metrix, len, col, row+pos, col-pos, pos, col-pos );
    for( i=pos;i<col;i++ )
    {
        for( j=row+pos;j<len;j++ )
        {
            metrix[i*len+j].poly2Int(inv);
            inv+=32;
        }
    }
    return (col-pos)*(col-pos);
}
*/
